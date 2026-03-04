package helper

import (
	"encoding/json"
	"net/http"
)

type Status string

const (
	StatusSuccess Status = "success"
	StatusError   Status = "error"
	StatusFailed  Status = "failed"
)

type ApiResponse struct {
	Code    int    `json:"code"`
	Status  Status `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func JsonResponse(w http.ResponseWriter, response *ApiResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.Code)
	json.NewEncoder(w).Encode(response)
}
