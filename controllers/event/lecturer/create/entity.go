package createEvent

type InputCreateEvent struct {
	LecturerId uint `json:"lecturer_id"`
	CourseId   uint `json:"course_id"`
	ClassId    uint `json:"class_id"`
	Meet       uint `json:"meet"`
}
