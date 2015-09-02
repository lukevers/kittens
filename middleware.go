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
