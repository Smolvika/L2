package handler

import (
	"develop/dev11"
	"net/http"
	"time"
)

func (h *Handler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	id, err := getUserId(r)
	if err != nil {
		WriteError(w, err)
		return
	}
	date, err := getDate(r)
	if err != nil {
		WriteError(w, err)
		return
	}

	ev, err := getInfo(r)
	if err != nil {
		WriteError(w, err)
		return
	}
	res, err := h.service.CreateEvent(dev11.EventQuery{
		UserId: id,
		Date:   date,
		Event:  ev,
	})

	if err != nil {
		WriteError(w, err)
		return
	}

	WriteResult(w, res)
}

func (h *Handler) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	id, err := getUserId(r)
	if err != nil {
		WriteError(w, err)
		return
	}

	ev, err := getInfo(r)
	if err != nil {
		WriteError(w, err)
		return
	}

	evId, err := getEventId(r)
	if err != nil {
		WriteError(w, err)
		return
	}

	res, err := h.service.UpdateEvent(dev11.EventAnswer{
		UserId:  id,
		EventId: evId,
		Event:   ev,
	})

	if err != nil {
		WriteError(w, err)
		return
	}

	WriteResult(w, res)
}

func (h *Handler) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	id, err := getUserId(r)
	if err != nil {
		WriteError(w, err)
		return
	}

	evId, err := getEventId(r)
	if err != nil {
		WriteError(w, err)
		return
	}

	res, err := h.service.DeleteEvent(id, evId)
	if err != nil {
		WriteError(w, err)
		return
	}
	WriteResult(w, res)
}

func (h *Handler) GetEventsForDay(w http.ResponseWriter, r *http.Request) {
	id, err := getUserId(r)
	if err != nil {
		WriteError(w, err)
		return
	}

	date, err := getDate(r)
	if err != nil {
		WriteError(w, err)
		return
	}

	res, err := h.service.GetEventsForDuration(date, id, 24*time.Hour)
	if err != nil {
		WriteError(w, err)
		return
	}

	WriteResultDuration(w, res)
}

func (h *Handler) GetEventsForWeek(w http.ResponseWriter, r *http.Request) {
	id, err := getUserId(r)
	if err != nil {
		WriteError(w, err)
		return
	}

	date, err := getDate(r)
	if err != nil {
		WriteError(w, err)
		return
	}

	res, err := h.service.GetEventsForDuration(date, id, 7*24*time.Hour)
	if err != nil {
		WriteError(w, err)
		return
	}

	WriteResultDuration(w, res)
}

func (h *Handler) GetEventsForMonth(w http.ResponseWriter, r *http.Request) {
	id, err := getUserId(r)
	if err != nil {
		WriteError(w, err)
		return
	}

	date, err := getDate(r)
	if err != nil {
		WriteError(w, err)
		return
	}

	res, err := h.service.GetEventsForDuration(date, id, 30*24*time.Hour)
	if err != nil {
		WriteError(w, err)
		return
	}

	WriteResultDuration(w, res)
}
