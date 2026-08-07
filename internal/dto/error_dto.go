package dto

import (
	"encoding/json"
	"time"
)

type ErrorResponse struct {
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
}

func NewErrorResponse(err error) ErrorResponse {
	return ErrorResponse{
		Message: err.Error(),
		Time:    time.Now(),
	}
}

func (e ErrorResponse) ToString() string {
	b, err := json.Marshal(e)
	if err != nil {
		return `{"message":"internal error"}`
	}
	return string(b)
}
