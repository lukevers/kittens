/* vim: set autoindent noexpandtab tabstop=4 shiftwidth=4: */
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tommy351/gin-sessions"
	"net/http"
)

func Guest() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Get(c)
		if session.Get("logged_in") == "true" {
			c.Redirect(http.StatusFound, "/")
			c.Abort()
		}

		c.Set("session", session)
		c.Next()
	}
}

func Authorized() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Get(c)

		if session.Get("logged_in") != "true" {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
		}

		// We need to make sure this user is actually logged in.
		id := session.Get("user_id")
		user := GetUser("id", id)
		// If the user id does not match, then we could not find the correct
		// user that they are claiming to be. A common example of when this
		// could happen is if someone manually updates the database.
		if user.ID != id {
			session.Clear()
			session.Save()
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
		}

		c.Next()
	}
}

func Expecting2Fa() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Get(c)

		if session.Get("needs_tfa") == "true" {
			c.Redirect(http.StatusFound, "/login/2fa")
			c.Abort()
		}
	}
}
