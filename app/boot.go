package app

import (
	"api/config"
	"api/model"
)

func Boot() {
	config.Load()
	config.ConnectDatabase()
	config.Database.AutoMigrate(&model.Users{})
	config.Database.AutoMigrate(&model.Transactions{})
}
