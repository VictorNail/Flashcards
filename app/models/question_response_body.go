package models

// QuestionResponseBody représente une réponse à une question.
type QuestionResponseBody struct {
	FlashcardID    int `json:"flashcardId"`
	NumeroResponse int `json:"numeroResponse"`
}
