package handler

import (
	"net/http"
	"schedule-api/internal/service"

	"github.com/gin-gonic/gin"
)

type ScheduleHandler struct {
	service *service.ScheduleService
}

func NewScheduleHander(s *service.ScheduleService) *ScheduleHandler {
	return &ScheduleHandler{service: s}
}

// GET /schedules
func (h *ScheduleHandler) GetSchedules(c *gin.Context) {
	schedules, err := h.service.GetSchedules()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, schedules)
}

// POST /schedules
func (h *ScheduleHandler) CreateSchedules(c *gin.Context) {
}

// func (s *service.ScheduleService) CreateSchedule(schedule *model.Schedule)

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Status": "ok",
	})
}
