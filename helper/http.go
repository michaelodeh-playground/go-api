package helper

import (
	"encoding/json"
	"net/http"
)

type Status string

const (
	StatusSuccess             Status = "success"
	StatusError               Status = "error"
	StatusFailed              Status = "failed"
	StatusNotFound            Status = "not_found"
	StatusBadRequest          Status = "bad_request"
	StatusInternalServerError Status = "internal_server_error"
)

type ApiResponse[T any] struct {
	Code    int    `json:"code"`
	Status  Status `json:"status"`
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`
}

type ApiSuccessResponse[T any] struct {
	Code    int    `json:"code" example:"200"`
	Status  Status `json:"status" example:"success"`
	Message string `json:"message" example:"Record fetched successfully"`
	Data    T      `json:"data"`
}

func JsonResponse[T any](w http.ResponseWriter, response *ApiResponse[T]) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.Code)
	json.NewEncoder(w).Encode(response)
}

func JsonErrorResponse[T any](w http.ResponseWriter, response *ApiResponse[T]) {
	JsonResponse(w, &ApiResponse[string]{
		Code:    response.Code,
		Status:  StatusError,
		Message: response.Message,
	})
}

func JsonNotFoundResponse(w http.ResponseWriter, message string) {
	JsonResponse(w, &ApiResponse[string]{
		Code:    http.StatusNotFound,
		Status:  StatusNotFound,
		Message: message,
	})
}

func JsonBadRequestResponse(w http.ResponseWriter, message string) {
	JsonResponse(w, &ApiResponse[string]{
		Code:    http.StatusBadRequest,
		Status:  StatusBadRequest,
		Message: message,
	})
}

func JsonInternalServerErrorResponse(w http.ResponseWriter, message string) {
	JsonResponse(w, &ApiResponse[string]{
		Code:    http.StatusInternalServerError,
		Status:  StatusInternalServerError,
		Message: message,
	})
}

func JsonSuccessResponse[T any](w http.ResponseWriter, response *ApiSuccessResponse[T]) {
	JsonResponse(w, &ApiResponse[T]{
		Code:    http.StatusOK,
		Status:  StatusSuccess,
		Message: response.Message,
		Data:    response.Data,
	})
}
