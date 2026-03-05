package wallet

import (
	"api/config"
	"api/helper"
	"api/model"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// @Summary Fund User Wallet
// @Description Fund user Wallet
// @Tags wallets
// @Accept json
// @Produce json
// @Param Request body FundWallet true "Fund Wallet Data"
// @Success 200 {object} helper.ApiSuccessResponse[model.Transactions]
// @Router /api/wallets/fund [post]
func Found(w http.ResponseWriter, r *http.Request) {
	var body FundWallet
	var user model.Users
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		helper.JsonBadRequestResponse(w, "Invalid JSON")
		return
	}

	if err := config.Database.First(&user, "id = ?", body.UserID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			helper.JsonNotFoundResponse(w, "User not found")
			return
		}
		helper.JsonInternalServerErrorResponse(w, "Internal Server Error")
		return
	}

	transaction := model.Transactions{
		Amount:      body.Amount,
		UserID:      user.ID,
		ReferenceID: uuid.New().String(),
	}
	err := config.Database.Transaction(func(trx *gorm.DB) error {
		if err := trx.Create(&transaction).Error; err != nil {
			return err
		}

		if err := trx.Model(&user).
			Where("id = ?", user.ID).
			Update("balance", gorm.Expr("balance + ?", body.Amount)).
			Error; err != nil {
			return err
		}

		// var u model.User
		// trx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&u, user.ID)

		return nil
	})

	if err != nil {
		helper.JsonInternalServerErrorResponse(w, "Transaction failed")
		return
	}
	config.Database.Preload("User").First(&transaction, "id = ?", transaction.ID)

	helper.JsonSuccessResponse(w, &helper.ApiSuccessResponse[model.Transactions]{
		Message: "Wallet funded successfully",
		Data:    transaction,
	})

}

// @Summary Get User Wallet Balance
// @Description Get user wallet balance
// @Tags wallets
// @Produce json
// @Param user path string true "User ID"
// @Success 200 {object} helper.ApiSuccessResponse[float64]
// @Router /api/wallets/balance/{user} [get]
func Balance(w http.ResponseWriter, r *http.Request) {
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

	helper.JsonSuccessResponse(w, &helper.ApiSuccessResponse[float64]{
		Message: "User fetched successfully",
		Data:    user.Balance,
	})

}
