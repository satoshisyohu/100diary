package domain

import (
	"github.com/satoshisyohu/pomodoro/pkg/domain/model"
	"gorm.io/gorm"
)

type TagRepository interface {
	Save(tx *gorm.DB, tag []*model.Tag) error
}
