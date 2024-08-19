package updateExcuse

type InputUpdateExcuse struct {
	ID     string
	Excuse string `json:"excuse"`
	Status uint   `json:"status"`
}
