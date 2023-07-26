package handler

import (
	"develop/dev11"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	InternalError      = "Internal Status Error"
	BadRequest         = "Status Bad Request"
	ServiceUnavailable = "Status Service Unavailable"
)

type CustomError struct {
	Code    int
	Message string
	ErrDesc string
}

func NewCustomError(code int, message string, errDesc string) *CustomError {
	return &CustomError{Code: code, Message: message, ErrDesc: errDesc}
}

func WriteError(w http.ResponseWriter, err *CustomError) {
	ans, _ := json.Marshal(map[string]interface{}{
		"error": *err,
	})
	fmt.Fprintf(w, "%s", string(ans))
}

func WriteResult(w http.ResponseWriter, res dev11.EventAnswer) {
	ans, _ := json.Marshal(map[string]interface{}{
		"result": res,
	})
	fmt.Fprintf(w, "%s", string(ans))
}
func WriteResultDuration(w http.ResponseWriter, res []dev11.Event) {
	ans, _ := json.Marshal(map[string]interface{}{
		"result": res,
	})
	fmt.Fprintf(w, "%s", string(ans))
}
