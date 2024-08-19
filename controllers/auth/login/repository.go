package loginAuth

import (
	model "attendance-is/models"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"gorm.io/gorm"
)

type Repository interface {
	LoginEmailAuthRepository(email string, password string) (*model.User, string)
	LoginNimAuthRepository(nim string, password string) (*model.User, string)
	LoginNipAuthRepository(nip string, password string) (*model.User, string)
}

type repository struct {
	db *gorm.DB
}

func NewLoginAuthRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) LoginEmailAuthRepository(email string, password string) (*model.User, string) {
	var user model.User
	if err := r.db.Model(user).Preload("Student").Preload("Lecturer").Preload("PBM").Where("email = ?", email).Take(&user).Error; err != nil {
		return nil, "LOGIN_UNAUTHENTICATED_401"
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return nil, "LOGIN_UNAUTHENTICATED_401"
	}
	return &user, ""
}

func (r *repository) LoginNimAuthRepository(nim string, password string) (*model.User, string) {
	var student model.Student
	r.db.Where("nim = ?", nim).Take(&student)
	var user model.User
	if err := r.db.Model(user).Preload("Student").Where("id = ?", student.UserID).Take(&user).Error; err != nil {
		return nil, err.Error()
	}

	fmt.Println(user.Student.Name)
	fmt.Println(user.Student.Nim)

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return nil, "LOGIN_UNAUTHENTICATED_401"
	}
	return &user, ""
}

func (r *repository) LoginNipAuthRepository(nip string, password string) (*model.User, string) {
	// var lecturer model.Lecturer
	// if r.db.Model(model.Lecturer{}).Where("nip = ?", nip).Preload("User").Find(&lecturer).Error != nil {
	// 	return nil, "LOGIN_UNAUTHENTICATED_401"
	// }
	// if bcrypt.CompareHashAndPassword([]byte(lecturer.User.Password), []byte(password)) != nil {
	// 	return nil, "LOGIN_UNAUTHENTICATED_401"
	// }
	// return &lecturer.User, ""
	return nil, ""
}
