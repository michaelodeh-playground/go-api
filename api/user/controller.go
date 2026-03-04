package user

import (
	httpStatusText "api/common"
	"api/config"
	"api/helper"
	"api/model"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// @Summary Create user
// @Description Create a new user
// @Tags users
// @Accept json
// @Produce json
// @Param user body CreateUserRequest true "Users Data"
// @Success 200 {object} model.Users
// @Router /api/users [post]
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var body CreateUserRequest

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		helper.JsonResponse(w, &helper.ApiResponse{
			Code:    http.StatusBadRequest,
			Status:  helper.StatusError,
			Message: "Invalid JSON",
		})
		return
	}
	if body.Name == "" {
		helper.JsonResponse(w, &helper.ApiResponse{
			Code:    http.StatusBadRequest,
			Status:  helper.StatusError,
			Message: "name is required",
		})
		return
	}

	user := model.Users{
		Name:  body.Name,
		Email: body.Email,
	}
	checkUser := config.Database.Find(&user, "email = ?", body.Email)

	if checkUser.RowsAffected > 0 {
		helper.JsonResponse(w, &helper.ApiResponse{
			Code:    http.StatusBadRequest,
			Status:  helper.StatusError,
			Message: "User already exists",
			Data:    user,
		})
		return
	}

	result := config.Database.Create(&user)
	if result.Error != nil {
		fmt.Println(result.Error)
		helper.JsonResponse(w, &helper.ApiResponse{
			Code:    http.StatusInternalServerError,
			Status:  helper.StatusError,
			Message: "Internal Server Error",
		})
		return
	}

	response := &helper.ApiResponse{
		Code:    http.StatusOK,
		Status:  httpStatusText.SUCCESS,
		Message: "User created successfully",
		Data:    user,
	}
	helper.JsonResponse(w, response)
}

// @Summary Get all users
// @Description Get all users
// @Tags users
// @Produce json
// @Success 200 {object} []model.Users
// @Router /api/users [get]
func IndexUser(w http.ResponseWriter, r *http.Request) {
	var users []model.Users
	result := config.Database.Find(&users)
	if result.Error != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := &helper.ApiResponse{
		Code:    http.StatusOK,
		Status:  httpStatusText.SUCCESS,
		Message: "User fetched successfully",
		Data:    users,
	}
	helper.JsonResponse(w, response)
}

// @Summary Update user
// @Description Update a user
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body UpdateUserRequest true "Users Data"
// @Success 200 {object} model.Users
// @Router /api/users/{user} [put]
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "user")
	var body UpdateUserRequest
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	user := config.Database.Where("id = ?", id)
	if user.RowsAffected == 0 {
		response := &helper.ApiResponse{
			Code:    http.StatusNotFound,
			Status:  httpStatusText.FAILED,
			Message: "User not found",
			Data:    nil,
		}
		helper.JsonResponse(w, response)
		return
	}
	result := user.Updates(&body)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	response := &helper.ApiResponse{
		Code:    http.StatusOK,
		Status:  httpStatusText.SUCCESS,
		Message: "User updated successfully",
		Data:    user,
	}
	helper.JsonResponse(w, response)
}
