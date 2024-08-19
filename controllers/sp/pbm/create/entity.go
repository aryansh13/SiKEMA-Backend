package createSP

type InputCreateSP struct {
	StudentID uint `json:"student_id"`
	Semester  uint `json:"semester"`
	Level     uint `json:"level"`
	Status    uint `json:"status"`
}
