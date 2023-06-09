package appresponse

import (
	"fmt"
	"net/http"
	"strings"
)

type Type string

type Response struct {
	ResponseCode int         `json:"responseCode" `
	Data         interface{} `json:"data,omitempty"`
	HttpCode     int         `json:"httpCode,omitempty"`
}

const (
	ErrorCode   = 4000
	SuccessCode = 1000
)

func NewAuthorizationError(reason interface{}) *Response {
	return &Response{
		ResponseCode: ErrorCode,
		Data:         reason,
		HttpCode:     http.StatusUnauthorized,
	}
}

func NewBadRequestError(reason interface{}) *Response {
	return &Response{
		ResponseCode: ErrorCode,
		Data:         reason,
		HttpCode:     http.StatusBadRequest,
	}
}

func NewConflictError(message interface{}) *Response {
	return &Response{
		ResponseCode: ErrorCode,
		Data:         message,
		HttpCode:     http.StatusConflict,
	}
}

func NewInternalError(msg ...string) *Response {
	message := strings.Join(msg, " ")
	return &Response{
		ResponseCode: ErrorCode,
		Data:         message,
		HttpCode:     http.StatusInternalServerError,
	}
}

func NewNotFoundError(message string) *Response {
	return &Response{
		ResponseCode: ErrorCode,
		Data:         message,
		HttpCode:     http.StatusNotFound,
	}
}

func NewPayloadTooLargeError(maxBodySize int64, contentLength int64) *Response {
	return &Response{
		ResponseCode: ErrorCode,
		Data:         fmt.Sprintf("Max payload size of %v exceeded. Actual payload size: %v", maxBodySize, contentLength),
		HttpCode:     http.StatusRequestEntityTooLarge,
	}
}

func NewServiceUnavailableError() *Response {
	return &Response{
		ResponseCode: ErrorCode,
		Data:         "Service unavailable or timed out",
		HttpCode:     http.StatusServiceUnavailable,
	}
}

func NewUnsupportedMediaTypeError(reason interface{}) *Response {
	return &Response{
		ResponseCode: ErrorCode,
		Data:         reason,
		HttpCode:     http.StatusUnsupportedMediaType,
	}
}

func NewSuccess(data interface{}) *Response {
	return &Response{
		ResponseCode: SuccessCode,
		Data:         data,
		HttpCode:     http.StatusOK,
	}
}
