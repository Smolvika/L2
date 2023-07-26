package service

import (
	"develop/dev11"
	"develop/dev11/pkg/handler"
	"develop/dev11/pkg/repository"
	"time"
)

type CalendarService struct {
	repo repository.Calendar
}

func NewCalendarService(repo repository.Calendar) *CalendarService {
	return &CalendarService{repo: repo}
}

func (s *CalendarService) CreateEvent(ev dev11.EventQuery) (dev11.EventAnswer, *handler.CustomError) {
	return s.repo.CreateEvent(ev)
}

func (s *CalendarService) UpdateEvent(ev dev11.EventAnswer) (dev11.EventAnswer, *handler.CustomError) {
	return s.repo.UpdateEvent(ev)
}

func (s *CalendarService) DeleteEvent(id, evId int) (dev11.EventAnswer, *handler.CustomError) {
	return s.repo.DeleteEvent(id, evId)
}

func (s CalendarService) GetEventsForDuration(date time.Time, id int, dur time.Duration) ([]dev11.Event, *handler.CustomError) {
	return s.repo.GetEventsForDuration(date, id, dur)
}
