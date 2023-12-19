package database

import (
	"os"

	"github.com/flutterninja9/shoppie/user_service/models"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DbInstance *gorm.DB

func InitializeDb(logger *logrus.Logger) error {
	logger.Info("Initializing database")
	db, err := gorm.Open(postgres.Open(os.Getenv("CONN_URL")), &gorm.Config{})
	if err != nil {
		logger.Fatal("Database initialization failed")
		return err
	}

	DbInstance = db
	migrationErr := DbInstance.AutoMigrate(&models.UserModel{})
	if migrationErr != nil {
		logger.Fatal("Unable to migrate database")
		return migrationErr
	}
	logger.Info("Database initialized sucessfully")

	return nil
}

func CreateUser(user *models.UserModel) (*models.UserModel, error) {
	result := DbInstance.Create(user)

	return user, result.Error
}

func GetUserById(id string) (*models.UserModel, error) {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	user := new(models.UserModel)

	result := DbInstance.First(user, "id = ?", uuid)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func GetUserByEmail(email string) (*models.UserModel, error) {
	user := new(models.UserModel)

	result := DbInstance.First(user, "email = ?", email)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
