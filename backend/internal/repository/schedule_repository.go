package repository

import (
	"schedule-api/internal/model"

	"gorm.io/gorm"
)

type ScheduleRepository struct {
	db *gorm.DB
}

func NewScheduleRepository(db *gorm.DB) *ScheduleRepository {
	return &ScheduleRepository{db: db}
}

func (r *ScheduleRepository) FindAll() ([]model.Schedule, error) {
	var schedules []model.Schedule
	err := r.db.Find(&schedules).Error
	return schedules, err
}

func (r *ScheduleRepository) Create(s *model.Schedule) error {
	return r.db.Create(s).Error
}

func (r *ScheduleRepository) Update(s *model.Schedule) error {
	return r.db.Save(s).Error
}

func (r *ScheduleRepository) Delete(id uint) error {
	return r.db.Delete(&model.Schedule{}, id).Error
}

func (r *ScheduleRepository) FindByID(id uint) (*model.Schedule, error) {
	var s model.Schedule
	err := r.db.First(&s, id).Error
	return &s, err
}
