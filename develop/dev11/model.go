package dev11

import "time"

//Событие

type Event struct {
	Date        time.Time `json:"time,omitempty"`
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
}

type EventQuery struct {
	UserId int
	Date   time.Time
	Event  Event
}

type EventAnswer struct {
	UserId  int
	EventId int
	Event   Event
}
