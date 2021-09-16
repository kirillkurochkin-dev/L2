package services

import (
	"dev11/internal/models"
	"time"
)

// Интерфейс сервиса, чтобы потом имплеменитровать его в db или in mem

type EventService interface {
	CreateEvent(event models.Event) error
	UpdateEvent(event models.Event) error
	DeleteEvent(id int) error
	GetAllEvents() ([]*models.Event, error)
	GetEventsForDay(userID int, date time.Time) ([]*models.Event, error)
	GetEventsForWeek(userID int, date time.Time) ([]*models.Event, error)
	GetEventsForMonth(userID int, date time.Time) ([]*models.Event, error)
}
