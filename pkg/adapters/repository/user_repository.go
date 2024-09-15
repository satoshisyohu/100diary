package repository

import (
	"github.com/satoshisyohu/pomodoro/pkg/domain"
	irepository "github.com/satoshisyohu/pomodoro/pkg/domain/repository"
	"gorm.io/gorm"
)

type userRepository struct{}

func NewUserRepository() irepository.UserRepository {
	return &userRepository{}
}

func (a *userRepository) Save(tx *gorm.DB, user *domain.User) error {
	if err := tx.Save(user).Error; err != nil {
		return err
	}
	return nil
}
