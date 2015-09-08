package main

import (
	"bufio"
	"bytes"
	"crypto/rand"
	"encoding/base32"
	"encoding/base64"
	"errors"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/dgryski/dgoogauth"
	"github.com/gin-gonic/gin"
	"github.com/tommy351/gin-sessions"
	"image/png"
	"net/http"
	"os"
)

func handleLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login", gin.H{
		"CAN_REGISTER": os.Getenv("USERS_CAN_REGISTER") == "true",
	})
}

func handleLoginPost(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	user := GetUser("username", username)

	if user.Id == 0 {
		// We don't want to give away too much information about what exactly
		// the error is here. If we say the username does not exist, then if
		// someone eventually hits a real username with a wrong password we
		// don't want to let them know they just need to figure out that
		// user's password. It's safer to respond with the same message
		// for both username and password errors.
		c.Error(errors.New("Could not authenticate"))
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"errors": c.Errors,
		})
	} else {
		if !user.AttemptPassword(password) {
			c.Error(errors.New("Could not authenticate"))
			c.JSON(http.StatusBadRequest, gin.H{
				"status": http.StatusBadRequest,
				"errors": c.Errors,
			})
		} else {
			session := sessions.Get(c)
			session.Set("logged_in", "true")
			session.Set("user_id", user.Id)

			if user.Twofa {
				session.Set("needs_tfa", "true")
			}

			session.Save()

			c.JSON(http.StatusOK, gin.H{
				"status": http.StatusOK,
				"twofa":  user.Twofa,
				"errors": c.Errors,
			})
		}
	}
}

func handle2faLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "2fa", nil)
}

func handle2faLoginPost(c *gin.Context) {
	token := c.PostForm("token")
	session := sessions.Get(c)
	user := GetUser("id", session.Get("user_id"))

	otpc := &dgoogauth.OTPConfig{
		Secret:      user.TwofaSecret,
		WindowSize:  3,
		HotpCounter: 0,
	}

	val, err := otpc.Authenticate(token)
	if err != nil || !val {
		c.Error(errors.New("Could not authenticate"))
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"errors": c.Errors,
		})
	} else {
		session.Set("needs_tfa", "false")
		session.Save()
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
		})
	}
}

func handleLogout(c *gin.Context) {
	session := sessions.Get(c)
	session.Clear()
	session.Save()

	c.Redirect(http.StatusFound, "/login")
}

func handleRegister(c *gin.Context) {
	c.HTML(http.StatusOK, "register", nil)
}

func handleRegisterPost(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	email := c.PostForm("email")

	if username == "" {
		c.Error(errors.New("Username cannot be blank"))
	}

	if password == "" {
		c.Error(errors.New("Password cannot be blank"))
	}

	if email == "" {
		c.Error(errors.New("Email cannot be blank"))
	}

	if len(c.Errors) > 0 {
		// If we have any errors, let's send them to the user now before
		// we do anything else. If we're missing any information, we
		// can't register properly anyways.
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"errors": c.Errors,
		})
	} else {
		// We need to make sure this username is not already taken
		user := GetUser("username", username)
		if user.Id != 0 {
			c.Error(errors.New("Username is already taken"))
			c.JSON(http.StatusBadRequest, gin.H{
				"status": http.StatusBadRequest,
				"errors": c.Errors,
			})
		} else {
			user = GetUser("email", email)
			if user.Id != 0 {
				c.Error(errors.New("Email is already taken"))
				c.JSON(http.StatusBadRequest, gin.H{
					"status": http.StatusBadRequest,
					"errors": c.Errors,
				})
			} else {
				// Now we can register!
				user = &User{
					Username: username,
					Email:    email,
				}

				// This function runs db.Save(&user), so we do not want to run
				// it again or we'll get a 1062 (duplicate entry) error.
				user.SetPassword(password)
				c.JSON(http.StatusOK, gin.H{
					"status": http.StatusOK,
					"errors": c.Errors,
				})
			}
		}
	}
}

func handleRoot(c *gin.Context) {
	c.HTML(http.StatusOK, "index", nil)
}

func handleSettings(c *gin.Context) {
	session := sessions.Get(c)
	user := GetUser("id", session.Get("user_id"))
	c.HTML(http.StatusOK, "settings", gin.H{
		"user": user,
	})
}

func handleSettingsUpdatePost(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	email := c.PostForm("email")

	session := sessions.Get(c)
	user := GetUser("id", session.Get("user_id"))

	// If they entered a password, let's update it.
	if password != "" {
		user.SetPassword(password)
	}

	if username == "" {
		// If the username is blank, they should get an error.
		c.Error(errors.New("Username can not be blank"))
	} else if user.Username != username {
		// If the username is not their current username, let's try
		// to update it.
		err := user.SetUsername(username)
		if err != nil {
			c.Error(err)
		}
	}

	if email == "" {
		// If the email is blank, they should get an error.
		c.Error(errors.New("Email can not be blank"))
	} else if user.Email != email {
		// If the email is not their current email, let's try to
		// update it.
		err := user.SetEmail(email)
		if err != nil {
			c.Error(err)
		}
	}

	if len(c.Errors) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"errors": c.Errors,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"errors": c.Errors,
		})
	}
}

func handleSettingsGenerate2fa(c *gin.Context) {
	session := sessions.Get(c)
	user := GetUser("id", session.Get("user_id"))

	// Get random secret
	s := make([]byte, 6)
	_, err := rand.Read(s)
	if err != nil {
		c.Error(errors.New("Could not generate random secret"))
	}

	secret := base32.StdEncoding.EncodeToString(s)
	session.Set("twofa_secret", secret)
	session.Save()
	// Create auth string to be encoded as a QR image
	//
	// https://github.com/google/google-authenticator/wiki/Key-Uri-Format
	// otpauth://totp/Example:alice@google.com?secret=JBSWY3DPEHPK3PXP&issuer=Example
	//
	authstr := "otpauth://totp/Kittens:" + user.Email + "?secret=" + secret + "&issuer=Kittens"

	// Encode the QR image
	qrcode, err := qr.Encode(authstr, qr.L, qr.Auto)
	if err != nil {
		c.Error(errors.New("Could not encode qr image"))
	}

	qrcode, err = barcode.Scale(qrcode, 512, 512)
	if err != nil {
		c.Error(errors.New("Could not scale qr image"))
	}

	var b bytes.Buffer
	buffer := bufio.NewWriter(&b)
	png.Encode(buffer, qrcode)
	buffer.Flush()

	data := base64.StdEncoding.EncodeToString(b.Bytes())

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"errors": c.Errors,
		"data":   data,
	})
}

func handleSettingsVerify2fa(c *gin.Context) {
	session := sessions.Get(c)
	user := GetUser("id", session.Get("user_id"))
	secret := session.Get("twofa_secret")

	token := c.PostForm("token")
	otpc := &dgoogauth.OTPConfig{
		Secret:      secret.(string),
		WindowSize:  3,
		HotpCounter: 0,
	}

	val, err := otpc.Authenticate(token)
	if err != nil {
		c.Error(errors.New("Could not authenticate token"))
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"errors": c.Errors,
		})
	} else if val {
		user.Twofa = true
		user.TwofaSecret = secret.(string)
		db.Save(&user)
		session.Delete("twofa_secret")
		session.Save()
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"errors": c.Errors,
		})
	}
}

func handleSettingsDisable2fa(c *gin.Context) {
	session := sessions.Get(c)
	user := GetUser("id", session.Get("user_id"))
	user.TwofaSecret = ""
	user.Twofa = false
	db.Save(&user)
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"errors": c.Errors,
	})
}
