package app

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/yenugantirahul/peersend/backend/internal/handlers"
	"github.com/yenugantirahul/peersend/backend/internal/routes"
	"github.com/yenugantirahul/peersend/backend/internal/services"
)

func New() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	sessionService := services.NewSessionService()
	sessionHandler := handlers.NewSessionHandler(sessionService)

	routes.Register(e, sessionHandler)

	return e
}
