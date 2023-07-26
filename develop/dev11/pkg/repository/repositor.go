package repository

import (
	"develop/dev11"
	"develop/dev11/pkg/handler"
	"time"
)

type Calendar interface {
	CreateEvent(ev dev11.EventQuery) (dev11.EventAnswer, *handler.CustomError)
	UpdateEvent(ev dev11.EventAnswer) (dev11.EventAnswer, *handler.CustomError)
	DeleteEvent(userId, eventId int) (dev11.EventAnswer, *handler.CustomError)
	GetEventsForDuration(date time.Time, userId int, dur time.Duration) ([]dev11.Event, *handler.CustomError)
}

type Repository struct {
	Calendar
}

func NewRepository() *Repository {
	return &Repository{Calendar: NewCache()}
}
