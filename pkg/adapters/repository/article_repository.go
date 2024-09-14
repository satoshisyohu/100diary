package repository

import (
	"github.com/satoshisyohu/pomodoro/pkg/domain"
	"github.com/satoshisyohu/pomodoro/pkg/domain/model"
	"gorm.io/gorm"
)

type articleRepository struct{}

func NewArticleRepository() domain.ArticleRepository {
	return &articleRepository{}
}

func (a *articleRepository) Save(tx *gorm.DB, article *model.Article) error {
	if err := tx.Save(article).Error; err != nil {
		return err
	}
	return nil
}
