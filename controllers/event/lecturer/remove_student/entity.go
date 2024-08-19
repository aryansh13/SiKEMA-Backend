package removeStudentEvent

type InputRemoveStudent struct {
	EventId   string
	StudentId []string `json:"student_id"`
}
