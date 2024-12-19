package session

import "Flashcards/app/services/session"

type Session struct {
	SessionService *session.Session
}

func New(sessionService *session.Session) *Session {
	return &Session{
		SessionService: sessionService,
	}
}
