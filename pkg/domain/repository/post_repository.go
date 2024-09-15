package repository

import (
	"github.com/satoshisyohu/pomodoro/pkg/domain"
	"gorm.io/gorm"
)

type PostRepository interface {
	Save(tx *gorm.DB, Post *domain.Post) error
	RetrieveByDateAndUserId(tx *gorm.DB, userId string) (*domain.Post, error)
}
