package handler

import (
	"encoding/json"
	"net/http"
)

const (
	ErrInvalidRequestBody  string = "invalid_request_body"
	ErrInternalServerError string = "internal_server_error"
)

type Error struct {
	Code         string `json:"code"`
	Message      string `json:"message"`
	MessageTitle string `json:"message_title"`
}

// Response is the structure for all API responses.
type Response struct {
	Data    interface{} `json:"data"`
	Success bool        `json:"success"`
	Errors  []Error     `json:"errors"`
}

// Marshal marshals the JSON response.
func (r *Response) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// NewErrorResponse creates a new error to be sent as the response.
func NewErrorResponse(code, message, title string) *Response {
	r := &Response{
		Data:    struct{}{},
		Success: false,
		Errors:  []Error{{code, message, title}},
	}
	return r
}

// NewSuccessResponse creates a new success response to be send via the API.
func NewSuccessResponse(data interface{}) *Response {
	r := &Response{
		Data:    data,
		Success: true,
		Errors:  []Error{},
	}
	return r
}

func (r *Response) Write(w http.ResponseWriter, status int) {
	body, err := r.Marshal()
	if err != nil {
		// add failure logs
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if _, err = w.Write(body); err != nil {
		// add logs
	}
}
