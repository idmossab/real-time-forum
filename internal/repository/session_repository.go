package repository

import (
	"database/sql"
	"fmt"
	"real_time_forum/internal/model"
	"time"
)

type SessionRepository interface {
	CreateSession(userID int, token string, expiresAt time.Time) error
	GetSessionByToken(token string) (*model.Session, error)
	DeleteSession(token string) error
	DeleteExpiredSessions() error
	DeleteUserSessions(userID int) error
}

type SessionRepo struct {
	DB *sql.DB
}

func (sr *SessionRepo) CreateSession(userID int, token string, expiresAt time.Time) error {
	// First, delete any existing sessions for this user (optional)
	_, err := sr.DB.Exec("DELETE FROM sessions WHERE user_id = ?", userID)
	if err != nil {
		return fmt.Errorf("error deleting existing sessions: %w", err)
	}

	// Create the new session
	_, err = sr.DB.Exec(
		"INSERT INTO sessions (user_id, session_token, expires_at) VALUES (?, ?, ?)",
		userID, token, expiresAt,
	)
	if err != nil {
		return fmt.Errorf("error creating session: %w", err)
	}

	return nil
}