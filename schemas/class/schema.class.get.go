package schemaClass

type GetClassResponse struct {
	ID       uint              `gorm:"primaryKey"`
	Name     string            `gorm:"type:varchar(25)"`
	Students []StudentResponse `gorm:"foreignKey:ClassID"`
}

type StudentResponse struct {
	ID   uint `gorm:"primaryKey"`
	Nim  string
	Name string
}

type GetStudentMeta struct {
	Count       int64
	Page        int64
	ItemPerPage int64
}

type CreateClassRequest struct {
	Name     string
	Students []uint
}

type UpdateClassRequest struct {
	ID       uint `json:"-"`
	Name     string
	Students []uint
}

type AddStudentToClassRequest struct {
	ID       uint
	Students []uint
}

type RemoveStudentFromClassRequest struct {
	ID      uint
	Student uint
}

type UpdateStudentFromClassRequest struct {
	ID       uint
	Students []uint
}
