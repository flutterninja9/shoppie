package database

import (
	"os"
	"product_service/models"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbInstance *gorm.DB

func InitDb(logger *logrus.Logger) error {
	logger.Info("Setting up DB")
	db, err := gorm.Open(postgres.Open(os.Getenv("CONN_URL")), &gorm.Config{})

	if err != nil {
		logger.Fatal("Unable to setup db")
		return err
	}

	dbInstance = db
	dbInstance.AutoMigrate(&models.ProductModel{})
	logger.Info("DB setup sucessfull")

	return nil
}

func CreateProduct(product *models.ProductModel) (*models.ProductModel, error) {
	result := dbInstance.Create(product)

	if result.Error != nil {
		return nil, result.Error
	}

	return product, nil
}

func GetProductById(id string) (*models.ProductModel, error) {
	product := new(models.ProductModel)
	result := dbInstance.First(product, "id=?", id)

	if result.Error != nil {
		return nil, result.Error
	}

	return product, nil
}

func GetAllProducts() ([]*models.ProductModel, error) {
	products := []*models.ProductModel{}
	result := dbInstance.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}
