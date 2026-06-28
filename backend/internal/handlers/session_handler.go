package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/yenugantirahul/peersend/backend/internal/services"
)

type SessionHandler struct {
	service *services.SessionService
}

func NewSessionHandler(service *services.SessionService) *SessionHandler {
	return &SessionHandler{
		service: service,
	}
}

func (h *SessionHandler) CreateSession(c echo.Context) error {
	session := h.service.CreateSession()
	return c.JSON(http.StatusCreated, session)
}
