package handler

import (
	"develop/dev11/pkg/service"
	"net/http"
)

type Handler struct {
	service *service.Service
}

func NewHandlers(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRouters() *http.Handler {
	eventMux := http.NewServeMux()
	eventMux.HandleFunc("/create_event", h.CreateEvent)
	eventMux.HandleFunc("/update_event", h.UpdateEvent)
	eventMux.HandleFunc("/delete_event", h.DeleteEvent)
	eventMux.HandleFunc("/events_for_day", h.GetEventsForDay)
	eventMux.HandleFunc("/events_for_week", h.GetEventsForWeek)
	eventMux.HandleFunc("/events_for_month", h.GetEventsForMonth)
	//Добавление Middleware
	siteHandler := LogMiddleware(eventMux)
	siteHandler = Recovery(siteHandler)
	return &siteHandler
}
