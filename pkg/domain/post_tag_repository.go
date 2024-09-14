package domain

import (
	"github.com/satoshisyohu/pomodoro/pkg/domain/model"
	"gorm.io/gorm"
)

type PostTagRepository interface {
	Save(tx *gorm.DB, postTag []*model.PostTag) error
}
