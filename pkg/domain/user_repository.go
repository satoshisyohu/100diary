package domain

import (
	"github.com/satoshisyohu/pomodoro/pkg/domain/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	Save(tx *gorm.DB, user *model.User) error
}
