package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tommy351/gin-csrf"
	"github.com/tommy351/gin-sessions"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	store  sessions.CookieStore
	router *gin.Engine
)

func InitRouter() {
	router = gin.New()
	store = sessions.NewCookieStore([]byte(os.Getenv("USERS_COOKIE_STORE_SECRET")))
	router.Use(gin.Logger())
	router.Use(sessions.Middleware("session", store))
	router.Use(csrf.Middleware(csrf.Options{
		Secret:    os.Getenv("USERS_CSRF_TOKEN_SECRET"),
		ErrorFunc: CsrfMismatch,
	}))

	if os.Getenv("APP_DEBUG") != "true" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Set custom HTML renderer so we can change the delimeters
	html, err := template.New("").Delims("[[", "]]").ParseGlob("assets/html/*.html")
	if err != nil {
		log.Fatal("Could not parse html files: ", err)
	}

	router.SetHTMLTemplate(html)

	addr := fmt.Sprintf("%s:%s",
		os.Getenv("WEB_INTERFACE"),
		os.Getenv("WEB_PORT"))

	server := &http.Server{
		Addr:           addr,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	router.Static("/assets", "./public/assets")
	router.StaticFile("/robots.txt", "./public/robots.txt")
	router.StaticFile("/favicon.ico", "./public/favicon.ico")

	AddRoutes()
	server.ListenAndServe()
}

func AddRoutes() {
	router.GET("/login", Guest(), handleLogin)
	router.POST("/login", Guest(), handleLoginPost)
	router.GET("/login/2fa", Authorized(), handle2faLogin)
	router.POST("/login/2fa", Authorized(), handle2faLoginPost)
	router.GET("/logout", Authorized(), handleLogout)

	if os.Getenv("USERS_CAN_REGISTER") == "true" {
		router.GET("/register", Guest(), handleRegister)
		router.POST("/register", Guest(), handleRegisterPost)
	}

	private := router.Group("/")
	private.Use(Authorized(), Expecting2Fa())
	{
		private.GET("/", handleRoot)
		private.GET("/settings", handleSettings)
		private.POST("/settings", handleSettingsUpdatePost)
	}
}

func CsrfMismatch(c *gin.Context) {
	c.String(400, "CSRF Token Mismatch")
}
