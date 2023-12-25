package database

import (
	"os"

	"github.com/flutterninja9/shoppie/cart_service/models"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DbInstance *gorm.DB

func InitDb(logger *logrus.Logger) {
	logger.Info("Initializing DB")

	db, err := gorm.Open(postgres.Open(os.Getenv("CONN_URL")), &gorm.Config{})

	if err != nil {
		logger.Panic(err)
	}

	DbInstance = db
	DbInstance.AutoMigrate(&models.CartModel{})
	DbInstance.AutoMigrate(&models.CartItem{})
	logger.Info("DB initialized")
}
