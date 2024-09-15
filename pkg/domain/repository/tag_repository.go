package repository

import (
	"github.com/satoshisyohu/pomodoro/pkg/domain"
	"gorm.io/gorm"
)

type TagRepository interface {
	Save(tx *gorm.DB, tag []*domain.Tag) error
}
