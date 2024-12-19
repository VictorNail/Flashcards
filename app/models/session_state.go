package models

// SessionState garde l'état d'avancement d'une session.
type SessionState struct {
	NextCardID int  `json:"nextCardId"`
	Score      int  `json:"score"`
	IsFinished bool `json:"isFinished"`
}
