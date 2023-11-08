package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"yoshikawa-keita-first-application/app/bff/handler"
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
