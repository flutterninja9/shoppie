package models

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;" json:"id"`
	Username  string    `gorm:"unique;not null" json:"username"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `gorm:"unique;not null" json:"email"`
	Password  string    `gorm:"not null" json:"-"`
}

func (user *UserModel) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.New()
	hashedPass, hasErr := hashPassword(user.Password)
	if hasErr != nil {
		return hasErr
	}

	user.Password = hashedPass
	return
}

func (user *UserModel) GenerateToken() (string, error) {
	token := jwt.New(jwt.SigningMethodHS512)
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}

	token.Claims = claims
	generated, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return "", nil
	}

	return generated, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func PasswordsMatch(password string, hashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))

	return err == nil
}
