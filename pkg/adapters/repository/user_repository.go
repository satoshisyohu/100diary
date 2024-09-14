package repository

import (
	"github.com/satoshisyohu/pomodoro/pkg/domain"
	"github.com/satoshisyohu/pomodoro/pkg/domain/model"
	"gorm.io/gorm"
)

type userRepository struct{}

func NewUserRepository() domain.UserRepository {
	return &userRepository{}
}

func (a *userRepository) Save(tx *gorm.DB, user *model.User) error {
	if err := tx.Save(user).Error; err != nil {
		return err
	}
	return nil
}
