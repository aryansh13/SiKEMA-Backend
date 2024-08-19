package service

import (
	model "attendance-is/models"

	schema "attendance-is/schemas/class"

	"gorm.io/gorm"
)

type ClassService struct {
	DB *gorm.DB
}

func NewClassService(db *gorm.DB) ClassService {
	return ClassService{DB: db}
}

func (s *ClassService) CreateClass(data schema.CreateClassRequest) (*model.Class, error) {
	var students []model.Student
	if err := s.DB.Where("id IN ?", data.Students).Find(&students).Error; err != nil {
		return nil, err
	}

	newClass := &model.Class{Name: data.Name, Students: students}
	if err := s.DB.Create(&newClass).Error; err != nil {
		return nil, err
	}

	return newClass, nil
}

func (s *ClassService) UpdateClass(data schema.UpdateClassRequest) (*model.Class, error) {
	var students []model.Student
	if err := s.DB.Where("ID IN (?)", data.Students).Find(&students).Error; err != nil {
		return nil, err
	}

	var class model.Class
	if err := s.DB.Where("id = ?", data.ID).Find(&class).Error; err != nil {
		return nil, err
	}

	if err := s.DB.Model(&class).Association("Students").Replace(&students); err != nil {
		return nil, err
	}

	class.Name = data.Name
	class.Students = students

	if err := s.DB.Save(&class).Error; err != nil {
		return nil, err
	}

	return &class, nil
}

func (s *ClassService) AddStudentToClass(data schema.AddStudentToClassRequest) (*model.Class, error) {
	var students []model.Student
	if err := s.DB.Where("id = ?", data.Students).Find(&students).Error; err != nil {
		return nil, err
	}

	var class model.Class
	if err := s.DB.Where("id = ?", data.ID).Find(&class).Error; err != nil {
		return nil, err
	}

	if err := s.DB.Model(&class).Association("Students").Append(&students); err != nil {
		return nil, err
	}

	return &class, nil
}

func (s *ClassService) RemoveStudentFromClass(data schema.RemoveStudentFromClassRequest) (*model.Class, error) {
	var class model.Class
	if err := s.DB.Where("id = ?", data.ID).Find(&class).Error; err != nil {
		return nil, err
	}

	var students []model.Student
	if err := s.DB.Where("id = ?", data.Student).Find(&students).Error; err != nil {
		return nil, err
	}

	if err := s.DB.Model(&class).Association("Students").Delete(&students); err != nil {
		return nil, err
	}

	return &class, nil
}

func (s *ClassService) UpdateStudentFromClass(data schema.UpdateStudentFromClassRequest) (*model.Class, error) {
	var class model.Class
	if err := s.DB.Where("id = ?", data.ID).Find(&class).Error; err != nil {
		return nil, err
	}

	var students []model.Student
	if err := s.DB.Where("id IN ?", data.Students).Find(&students).Error; err != nil {
		return nil, err
	}

	if err := s.DB.Model(&class).Association("Students").Replace(&students); err != nil {
		return nil, err
	}

	return &class, nil
}

func (s *ClassService) GetClass(meta *schema.GetStudentMeta) (*[]model.Class, error) {
	var class []model.Class
	if err := s.DB.Model(&model.Class{}).Count(&meta.Count).Limit(int(meta.ItemPerPage)).Offset(int(meta.ItemPerPage * (meta.Page - 1))).Preload("Students").Find(&class).Error; err != nil {
		return nil, err
	}

	return &class, nil
}

func (s *ClassService) ShowClass(ID string) (*model.Class, error) {
	var class model.Class

	if err := s.DB.Preload("Students").Where("id = ?", ID).Find(&class).Error; err != nil {
		return nil, err
	}

	return &class, nil
}
