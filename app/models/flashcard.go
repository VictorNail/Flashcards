package models

// Flashcard représente une carte de type question/réponse.
type Flashcard struct {
	Answer           string         `json:"answer"`
	Responses        []ResponseCard `json:"responses"`
	NumRightResponse int            `json:"numRightResponse"`
	Tags             []string       `json:"tags"`
}

// ResponseCard représente une proposition de réponse.
type ResponseCard struct {
	ID       int    `json:"id"`       // ID entre 1 et 4
	Proposal string `json:"proposal"` // Proposition de réponse
}
