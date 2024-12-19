package flashcard

import "Flashcards/app/services/flashcard"

type Flashcard struct {
	FlashcardService *flashcard.Flashcard
}

func New(flashcardService *flashcard.Flashcard) *Flashcard {
	return &Flashcard{
		FlashcardService: flashcardService,
	}
}
