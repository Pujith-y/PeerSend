// routes.go
package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/yenugantirahul/peersend/backend/internal/handlers"
)

func Register(e *echo.Echo, sessionHandler *handlers.SessionHandler) {
	e.POST("/sessions", sessionHandler.CreateSession)
}
