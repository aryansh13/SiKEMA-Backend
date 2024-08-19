package addStudentEvent

type InputAddStudent struct {
	EventId   string
	StudentId []string `json:"student_id"`
}
