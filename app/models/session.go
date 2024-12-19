package models

// Session représente une session utilisateur.
type Session struct {
	StudentID     string         `json:"studentID"`
	SessionID     string         `json:"sessionID"`
	Score         int            `json:"score"`
	Category      string         `json:"category"`
	FlashcardList []Flashcard    `json:"flashcardList"`
	ProposalList  []ResponseCard `json:"proposalList"`
	IsFinished    bool           `json:"isFinished"`
}
