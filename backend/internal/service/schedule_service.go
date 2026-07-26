package service

import (
	"schedule-api/internal/model"
	"schedule-api/internal/repository"
)

type ScheduleService struct {
	repo *repository.ScheduleRepository
}

func NewScheduleService(repo *repository.ScheduleRepository) *ScheduleService {
	return &ScheduleService{repo: repo}
}

func (s *ScheduleService) GetSchedules() ([]model.Schedule, error) {
	return s.repo.FindAll()
}

func (s *ScheduleService) CreateSchedule(sc *model.Schedule) error {
	return s.repo.Create(sc)
}

func (s *ScheduleService) UpdateSchedule(sc *model.Schedule) error {
	return s.repo.Update(sc)
}

func (s *ScheduleService) DeleteSchedule(id uint) error {
	return s.repo.Delete(id)
}
