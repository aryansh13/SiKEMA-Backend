package updateEvent

import (
	model "attendance-is/models"

	"gorm.io/gorm"
)

type Repository interface {
	GetEventByID(ID string) (*model.Event, string)
	SaveEvent(event *model.Event) string
}

type repository struct {
	db *gorm.DB
}

func NewUpdateEventRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) GetEventByID(ID string) (*model.Event, string) {
	var event model.Event
	if err := r.db.Model(&event).Preload("Course").Preload("Students").Where("id = ?", ID).Take(&event).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, "EVENT_NOTFOUND_404"
		}
		return nil, "EVENT_UNEXPECTED_500 : " + err.Error()
	}
	return &event, ""
}

func (r *repository) SaveEvent(event *model.Event) string {
	if err := r.db.Save(event).Error; err != nil {
		if err == gorm.ErrDuplicatedKey {
			return "ABSENT_CONFLICT_409"
		}
		return "ABSENT_UNEXPECTED_500 : " + err.Error()
	}

	return ""
}
