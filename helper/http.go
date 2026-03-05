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

type ApiResponse struct {
	Code    int    `json:"code"`
	Status  Status `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type ApiSuccessResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func JsonResponse(w http.ResponseWriter, response *ApiResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.Code)
	json.NewEncoder(w).Encode(response)
}

func JsonErrorResponse(w http.ResponseWriter, response *ApiResponse) {
	JsonResponse(w, &ApiResponse{
		Code:    response.Code,
		Status:  StatusError,
		Message: response.Message,
	})
}

func JsonNotFoundResponse(w http.ResponseWriter, message string) {
	JsonResponse(w, &ApiResponse{
		Code:    http.StatusNotFound,
		Status:  StatusNotFound,
		Message: message,
	})
}

func JsonBadRequestResponse(w http.ResponseWriter, message string) {
	JsonResponse(w, &ApiResponse{
		Code:    http.StatusBadRequest,
		Status:  StatusBadRequest,
		Message: message,
	})
}

func JsonInternalServerErrorResponse(w http.ResponseWriter, message string) {
	JsonResponse(w, &ApiResponse{
		Code:    http.StatusInternalServerError,
		Status:  StatusInternalServerError,
		Message: message,
	})
}

func JsonSuccessResponse(w http.ResponseWriter, response *ApiSuccessResponse) {
	JsonResponse(w, &ApiResponse{
		Code:    http.StatusOK,
		Status:  StatusSuccess,
		Message: response.Message,
		Data:    response.Data,
	})
}
