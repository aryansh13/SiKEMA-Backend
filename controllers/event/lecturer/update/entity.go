package updateEvent

type InputUpdateEvent struct {
	EventID string `json:"event_id"`
	Status  uint   `json:"status"`
}
