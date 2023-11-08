package router

import (
	"github.com/Yoshikawa-Keita/first-application/app/bff/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter() *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:8081"},
		AllowHeaders:     []string{"*"},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowCredentials: true,
	}))
	e.GET("/hello", handler.Hello)
	e.GET("/health", handler.Health)
	return e
}
