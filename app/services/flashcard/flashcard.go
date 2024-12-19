package flashcard

import "github.com/go-playground/validator/v10"

type Flashcard struct {
	validate *validator.Validate
}

func New() *Flashcard {
	return &Flashcard{
		validate: validator.New(),
	}
}
