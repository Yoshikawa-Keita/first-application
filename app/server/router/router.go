package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	serverHandler "yoshikawa-keita-first-application/app/server/handler"
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
