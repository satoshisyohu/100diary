package domain

import (
	"github.com/satoshisyohu/pomodoro/pkg/domain/model"
	"gorm.io/gorm"
)

type ArticleRepository interface {
	Save(tx *gorm.DB, article *model.Article) error
}
