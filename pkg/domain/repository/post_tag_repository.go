package repository

import (
	"github.com/satoshisyohu/pomodoro/pkg/domain"
	"gorm.io/gorm"
)

type PostTagRepository interface {
	Save(tx *gorm.DB, postTag []*domain.PostTag) error
}
