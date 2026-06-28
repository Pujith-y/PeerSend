package services

import (
	"time"

	"github.com/google/uuid"
	"github.com/yenugantirahul/peersend/backend/internal/model"
)

type SessionService struct{}

func NewSessionService() *SessionService {
	return &SessionService{}
}

func (s *SessionService) CreateSession() models.Session {
	return models.Session{
		ID:        uuid.NewString(),
		ExpiresAt: time.Now().Add(5 * time.Minute),
	}
}
