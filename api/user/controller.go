package user

import (
	httpStatusText "api/common"
	"api/config"
	"api/helper"
	"api/model"
	"encoding/json"
	"net/http"
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
	var user CreateUserRequest

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	config.Database.Create(&model.Users{
		Name:  user.Name,
		Email: user.Email,
		Age:   user.Age,
	})

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

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user model.Users

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	response := &helper.ApiResponse{
		Code:    http.StatusOK,
		Status:  helper.StatusSuccess,
		Message: "User updated successfully",
		Data:    user,
	}
	helper.JsonResponse(w, response)
}
