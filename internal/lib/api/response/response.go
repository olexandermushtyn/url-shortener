package response

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
	Alias  string `json:"alias,omitempty"`
}

const (
	StatusOK    = "OK"
	StatusError = "ERROR"
)

func OK() Response {
	return Response{Status: StatusOK}
}

func Error(msg string) Response {
	return Response{Status: StatusError, Error: msg}
}

func ValidateError(errs validator.ValidationErrors) Response {
	var errMsgs []string
	
	for _, err := range errs {
		switch err.Tag() {
		case "required":
			errMsgs = append(errMsgs, fmt.Sprintf("%s is required", err.Field()))
		case "url":
			errMsgs = append(errMsgs, fmt.Sprintf("%s is not a valid URL", err.Field()))
		default:
			errMsgs = append(errMsgs, fmt.Sprintf("%s is invalid", err.Field()))
		}
	}

	return Response{Status: StatusError, Error: strings.Join(errMsgs, ", ")}
}
