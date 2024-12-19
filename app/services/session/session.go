package session

import "github.com/go-playground/validator/v10"

type Session struct {
	validate *validator.Validate
}

func New() *Session {
	return &Session{
		validate: validator.New(),
	}
}
