package appresponse

import (
	"fmt"
	"net/http"
	"strings"
)

type Type string

type Success struct {
	ResponseCode uint        `json:"responseCode"`
	Message      Type        `json:"message,omitempty"`
	Data         interface{} `json:"data,omitempty"`
	HttpCode     int         `json:"httpCode,omitempty"`
}

type Error struct {
	ResponseCode uint        `json:"responseCode"`
	Message      Type        `json:"message,omitempty"`
	Error        interface{} `json:"error,omitempty"`
	HttpCode     int         `json:"httpCode,omitempty"`
}

const (
	Authorization        Type = "AUTHORIZATION"
	BadRequest           Type = "BAD_REQUEST"
	Conflict             Type = "CONFLICT"
	Internal             Type = "INTERNAL"
	NotFound             Type = "NOT_FOUND"
	PayloadTooLarge      Type = "PAYLOAD_TOO_LARGE"
	ServiceUnavailable   Type = "SERVICE_UNAVAILABLE"
	UnsupportedMediaType Type = "UNSUPPORTED_MEDIA_TYPE"
)

const (
	ErrorCode   = uint(4000)
	SuccessCode = uint(1000)
)

func NewAuthorizationError(reason interface{}) *Error {
	return &Error{
		ResponseCode: ErrorCode,
		Message:      Authorization,
		Error:        reason,
		HttpCode:     http.StatusUnauthorized,
	}
}

func NewBadRequestError(reason interface{}) *Error {
	return &Error{
		ResponseCode: ErrorCode,
		Message:      BadRequest,
		Error:        reason,
		HttpCode:     http.StatusBadRequest,
	}
}

func NewConflictError(message interface{}) *Error {
	return &Error{
		ResponseCode: ErrorCode,
		Message:      Conflict,
		Error:        message,
		HttpCode:     http.StatusConflict,
	}
}

func NewInternalError(msg ...string) *Error {
	message := strings.Join(msg, " ")
	return &Error{
		ResponseCode: ErrorCode,
		Message:      Internal,
		Error:        message,
		HttpCode:     http.StatusInternalServerError,
	}
}

func NewNotFoundError(message string) *Error {
	return &Error{
		ResponseCode: ErrorCode,
		Message:      NotFound,
		Error:        message,
		HttpCode:     http.StatusNotFound,
	}
}

func NewPayloadTooLargeError(maxBodySize int64, contentLength int64) *Error {
	return &Error{
		ResponseCode: ErrorCode,
		Message:      PayloadTooLarge,
		Error:        fmt.Sprintf("Max payload size of %v exceeded. Actual payload size: %v", maxBodySize, contentLength),
		HttpCode:     http.StatusRequestEntityTooLarge,
	}
}

func NewServiceUnavailableError() *Error {
	return &Error{
		ResponseCode: ErrorCode,
		Message:      ServiceUnavailable,
		Error:        "Service unavailable or timed out",
		HttpCode:     http.StatusServiceUnavailable,
	}
}

func NewUnsupportedMediaTypeError(reason interface{}) *Error {
	return &Error{
		ResponseCode: ErrorCode,
		Message:      UnsupportedMediaType,
		Error:        reason,
		HttpCode:     http.StatusUnsupportedMediaType,
	}
}

func NewSuccess(data interface{}) *Success {
	return &Success{
		ResponseCode: SuccessCode,
		Message:      "OK",
		Data:         data,
		HttpCode:     http.StatusOK,
	}
}
