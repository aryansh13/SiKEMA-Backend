package createExcuse

type InputCreateExcuse struct {
	ID         uint   `json:"id"`
	AbsentID   uint   `json:"absent_id"`
	Excuse     string `json:"excuse"`
	Attachment string `json:"attachment"`
}
