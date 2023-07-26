package handler

import (
	"develop/dev11"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("PanicMiddleWare: %s\n", r.URL.Path)
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("recovered", err)
				WriteError(w, &CustomError{
					Code:    500,
					Message: InternalError,
					ErrDesc: fmt.Sprintf("%s", err),
				})
				return
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("LogMiddleware: %s\n", r.URL.Path)
		start := time.Now()
		next.ServeHTTP(w, r)
		fmt.Printf("[%s] %s, %s %s\n",
			r.Method, r.RemoteAddr, r.URL.Path, time.Since(start))
	})
}

func getUserId(r *http.Request) (int, *CustomError) {
	//Получение id user
	id := r.URL.Query().Get("user_id")
	//Проверка id user
	userId, err := strconv.Atoi(id)
	if err != nil {
		return 0, NewCustomError(400, BadRequest, err.Error())
	}
	return userId, nil
}

func getDate(r *http.Request) (time.Time, *CustomError) {
	//Получение date
	date := r.URL.Query().Get("date")
	//Проверка date
	dateT, err := time.Parse("2006-01-02", date)

	if err != nil {
		return time.Time{}, NewCustomError(400, BadRequest, err.Error())
	}
	return dateT, nil
}

func getInfo(r *http.Request) (dev11.Event, *CustomError) {
	//Получение информации о событии
	buf, err := io.ReadAll(r.Body)
	if err != nil {
		return dev11.Event{}, NewCustomError(500, InternalError, err.Error())
	}
	var ev dev11.Event
	err = json.Unmarshal(buf, &ev)

	if err != nil {
		return dev11.Event{}, NewCustomError(400, BadRequest, err.Error())
	}
	return ev, nil
}

func getEventId(r *http.Request) (int, *CustomError) {
	//Получение id event
	id := r.URL.Query().Get("event_id")
	//Проверка id event
	eventId, err := strconv.Atoi(id)
	if err != nil {
		return 0, NewCustomError(400, BadRequest, err.Error())
	}
	return eventId, nil
}
