package router

import (
	serverHandler "github.com/Yoshikawa-Keita/first-application/app/server/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter() *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:8080"},
		AllowHeaders:     []string{"*"},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowCredentials: true,
	}))
	e.POST("/fizzbuzz", serverHandler.FizzBuzz)
	e.GET("/health", serverHandler.Health)
	return e
}
