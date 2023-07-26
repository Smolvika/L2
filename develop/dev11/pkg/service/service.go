package service

import (
	"develop/dev11"
	"develop/dev11/pkg/handler"
	"develop/dev11/pkg/repository"
	"time"
)

type Calendar interface {
	CreateEvent(ev dev11.EventQuery) (dev11.EventAnswer, *handler.CustomError)
	UpdateEvent(buf dev11.EventAnswer) (dev11.EventAnswer, *handler.CustomError)
	DeleteEvent(id, evId int) (dev11.EventAnswer, *handler.CustomError)
	GetEventsForDuration(date time.Time, id int, dur time.Duration) ([]dev11.Event, *handler.CustomError)
}

type Service struct {
	Calendar
}

func NewService(repos *repository.Repository) *Service {
	return &Service{Calendar: NewCalendarService(repos.Calendar)}
}
