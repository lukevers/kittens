package main

import (
	"github.com/dgryski/dgoogauth"
	"github.com/gin-gonic/gin"
	"github.com/tommy351/gin-sessions"
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
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		if !user.AttemptPassword(password) {
			c.AbortWithStatus(http.StatusBadRequest)
		} else {
			session := sessions.Get(c)
			session.Set("logged_in", "true")
			session.Set("user_id", user.Id)

			if user.Twofa {
				session.Set("needs_tfa", "true")
			}

			session.Save()

			c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "twofa": user.Twofa})
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

	// TODO
	// - this should work, but I need to test it later

	val, err := otpc.Authenticate(token)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else if val {
		session.Set("needs_tfa", "false")
		session.Save()
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
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
	//password := c.PostForm("password")

	// We need to make sure this username is not already taken
	user := GetUser("username", username)
	if user.Id != 0 {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		// TODO
	}
}

func handleRoot(c *gin.Context) {
	c.HTML(http.StatusOK, "index", nil)
}

func handleSettings(c *gin.Context) {
	c.HTML(http.StatusOK, "settings", nil)
}
