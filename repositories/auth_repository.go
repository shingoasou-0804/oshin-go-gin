package repositories

import (
	"errors"

	"github.com/shingoasou-0804/oshin-go-gin/models"
	"gorm.io/gorm"
)

type IAuthRepository interface {
	CreateUser(user models.User) error
	FindUserByEmail(email string) (*models.User, error)
}

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) IAuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) CreateUser(user models.User) error {
	result := r.db.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *AuthRepository) FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	result := r.db.First(&user, "email = ?", email)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return nil, errors.New("User not found")
		}
		return nil, result.Error
	}
	return &user, nil
}
