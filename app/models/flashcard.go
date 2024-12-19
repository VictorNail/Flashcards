package models

import "time"

// Flashcard représente une carte de type question/réponse.
type Flashcard struct {
	CustomID         string         `bson:"customID" json:"customId"`
	CreatedAt        time.Time      `json:"creationDate"`
	Answer           string         `json:"answer"`
	Responses        []ResponseCard `json:"responses"`
	NumRightResponse int            `json:"numRightResponse"`
	Tags             []string       `json:"tags"`
}

func (f *Flashcard) Collection() string {
	return "flashcards"
}

// ResponseCard représente une proposition de réponse.
type ResponseCard struct {
	ID       int    `json:"id"`       // ID entre 1 et 4
	Proposal string `json:"proposal"` // Proposition de réponse
}

// Flashcardinput
type FlashcardInput struct {
	Answer           string         `json:"answer"`
	Responses        []ResponseCard `json:"responses"`
	NumRightResponse int            `json:"numRightResponse"`
	Tags             []string       `json:"tags"`
}
