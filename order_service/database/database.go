package database

import (
	"os"

	"github.com/flutterninja9/shoppie/order_service/models"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DbInstance *gorm.DB

func InitDb(logger *logrus.Logger) error {
	logger.Info("Setting up DB")
	db, err := gorm.Open(postgres.Open(os.Getenv("CONN_URL")), &gorm.Config{})

	if err != nil {
		logger.Fatal("Unable to setup db")
		return err
	}

	DbInstance = db
	DbInstance.AutoMigrate(&models.OrderItem{})
	DbInstance.AutoMigrate(&models.OrderModel{})
	logger.Info("DB setup sucessfull")

	return nil
}
