package user

import (
	"api/config"
	"api/helper"
	"api/model"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

// @Summary Get all users
// @Description Get all users
// @Tags users
// @Produce json
// @Success 200 {object} helper.ApiSuccessResponse[[]model.Users]
// @Router /api/users [get]
func Index(w http.ResponseWriter, r *http.Request) {
	var users []model.Users
	result := config.Database.Find(&users)
	if result.Error != nil {
		helper.JsonInternalServerErrorResponse(w, "Internal Server Error")
		return
	}

	helper.JsonSuccessResponse(w, &helper.ApiSuccessResponse[[]model.Users]{
		Message: "User fetched successfully",
		Data:    users,
	})
}

// @Summary Show user
// @Description Show a user
// @Tags users
// @Produce json
// @Param user path string true "User ID"
// @Success 200 {object} helper.ApiSuccessResponse[model.Users]
// @Router /api/users/{user} [get]
func Show(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "user")
	var user model.Users
	if err := config.Database.First(&user, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			helper.JsonNotFoundResponse(w, "User not found")
			return
		}
		helper.JsonInternalServerErrorResponse(w, "Internal Server Error")
		return
	}
	helper.JsonSuccessResponse(w, &helper.ApiSuccessResponse[model.Users]{
		Message: "User fetched successfully",
		Data:    user,
	})
}

// @Summary Create user
// @Description Create a new user
// @Tags users
// @Accept json
// @Produce json
// @Param user body CreateUserRequest true "Users Data"
// @Success 200 {object} helper.ApiSuccessResponse[model.Users]
// @Router /api/users [post]
func Create(w http.ResponseWriter, r *http.Request) {
	var body CreateUserRequest

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		helper.JsonBadRequestResponse(w, "Invalid JSON")
		return
	}
	if body.Name == "" {
		helper.JsonBadRequestResponse(w, "name is required")
		return
	}
	if body.Email == "" {
		helper.JsonBadRequestResponse(w, "email is required")
		return
	}

	user := model.Users{
		Name:  body.Name,
		Email: body.Email,
	}
	var existing model.Users
	if err := config.Database.Where("email = ?", body.Email).First(&existing).Error; err == nil {
		helper.JsonBadRequestResponse(w, "User already exists")
		return
	}
	result := config.Database.Create(&user)
	if result.Error != nil {
		fmt.Println(result.Error)
		helper.JsonInternalServerErrorResponse(w, "Internal Server Error")
		return
	}

	helper.JsonSuccessResponse(w, &helper.ApiSuccessResponse[model.Users]{
		Message: "User created successfully",
		Data:    user,
	})
}

// @Summary Update user
// @Description Update a user
// @Tags users
// @Accept json
// @Produce json
// @Param user path string true "User ID"
// @Param user body UpdateUserRequest true "Users Data"
// @Success 200 {object} helper.ApiSuccessResponse[model.Users]
// @Router /api/users/{user} [put]
func Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "user")
	var body UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		helper.JsonBadRequestResponse(w, "Invalid JSON")
		return
	}
	var user model.Users
	if err := config.Database.First(&user, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			helper.JsonNotFoundResponse(w, "User not found")
			return
		}
		helper.JsonInternalServerErrorResponse(w, "Internal Server Error")
		return
	}

	if err := config.Database.Model(&user).Updates(body).Error; err != nil {
		helper.JsonInternalServerErrorResponse(w, "Internal Server Error")
		return
	}
	helper.JsonSuccessResponse(w, &helper.ApiSuccessResponse[model.Users]{
		Message: "User updated successfully",
		Data:    user,
	})

}

// @Summary Delete user
// @Description Delete a user
// @Tags users
// @Produce json
// @Param user path string true "User ID"
// @Success 200 {object} helper.ApiSuccessResponse[model.Users]
// @Router /api/users/{user} [delete]
func Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "user")
	var user model.Users
	if err := config.Database.First(&user, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			helper.JsonNotFoundResponse(w, "User not found")
			return
		}
		helper.JsonInternalServerErrorResponse(w, "Internal Server Error")
		return
	}
	result := config.Database.Delete(&user)
	if result.Error != nil {
		helper.JsonInternalServerErrorResponse(w, "Internal Server Error")
		return
	}
	if result.RowsAffected == 0 {
		helper.JsonErrorResponse(w, &helper.ApiResponse[string]{
			Message: "User not deleted",
		})
		return
	}

	helper.JsonSuccessResponse(w, &helper.ApiSuccessResponse[model.Users]{
		Message: "User deleted successfully",
		Data:    user,
	})
}
