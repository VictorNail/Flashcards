package models

import "time"

// Session représente une session utilisateur.
type Session struct {
	CustomID      string         `bson:"customID" json:"customId"`
	StudentID     string         `json:"studentID"`
	SessionID     string         `json:"sessionID"`
	Score         int            `json:"score"`
	Category      string         `json:"category"`
	FlashcardList []Flashcard    `json:"flashcardList"`
	ProposalList  []ResponseCard `json:"proposalList"`
	IsFinished    bool           `json:"isFinished"`
	CreatedAt     time.Time      `json:"creationDate"`
	Suspended     bool           `json:"suspended"`
}

// Collection Mongodb collection
func (s *Session) Collection() string {
	return "session"
}

// SessionInput
type SessionInput struct {
	StudentID     string         `json:"studentID"`
	SessionID     string         `json:"sessionID"`
	Score         int            `json:"score"`
	Category      string         `json:"category"`
	FlashcardList []Flashcard    `json:"flashcardList"`
	ProposalList  []ResponseCard `json:"proposalList"`
	IsFinished    bool           `json:"isFinished"`
}
