package main

import (
	"encoding/json"
	"net/http"
)

type BaseResponse struct {
	Status  string       `json:"status"`
	Message string       `json:"message"`
	Data    *interface{} `json:"data,omitempty"`
	Error   *ErrorDetail `json:"error,omitempty"`
}

type ErrorDetail struct {
	Code    int    `json:"code"`
	Details string `json:"details"`
}

func (app *application) successResponse(w http.ResponseWriter, code int, message string, data interface{}) {
	resp := BaseResponse{
		Status:  "success",
		Message: message,
		Data:    &data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)
}

func (app *application) failedResponse(w http.ResponseWriter, code int, message string, err error) {
	resp := BaseResponse{
		Status:  "error",
		Message: message,
		Error: &ErrorDetail{
			Code:    code,
			Details: err.Error(),
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)
}
