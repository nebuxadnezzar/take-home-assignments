package model

import (
	"fmt"
	"net/http"
)

type ErrMsg struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

func NewErrorMessage(statusCode int, message string, err error) ErrMsg {
	if err == nil {
		err = fmt.Errorf("unknown error")
	}
	errorMessage := ErrMsg{
		Status:  fmt.Sprintf("%v %s", statusCode, http.StatusText(statusCode)),
		Message: message,
		Error:   err.Error(),
	}

	return errorMessage
}
