package updateExcuse

type InputUpdateExcuse struct {
	ID         string
	Excuse     string `json:"excuse"`
	Attachment string `json:"attachment"`
}
