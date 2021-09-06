package main

import (
	"net/http"
	"github.com/labstack/echo"
)

func main() {
	port := getenv("PORT", "8080")

	e := echo.New()
	h := &handler{}

	e.GET("/", h.index)
	e.POST("/login", h.login)
	e.POST("/predict", h.predict, isLoggedIn)
	e.GET("/admin", h.private, isLoggedIn, isAdmin)
	e.POST("/token", h.token)

	e.Logger.Fatal(e.Start(":" + port))
}
