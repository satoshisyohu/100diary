package repository

import (
	"github.com/satoshisyohu/pomodoro/pkg/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	Save(tx *gorm.DB, user *domain.User) error
}
