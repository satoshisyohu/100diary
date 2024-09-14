package repository

import (
	"github.com/satoshisyohu/pomodoro/pkg/domain"
	"github.com/satoshisyohu/pomodoro/pkg/domain/model"
	"gorm.io/gorm"
)

type postTagRepository struct{}

func NewPostTagRepository() domain.PostTagRepository {
	return &postTagRepository{}
}

func (a *postTagRepository) Save(tx *gorm.DB, postTags []*model.PostTag) error {
	if err := tx.Save(postTags).Error; err != nil {
		return err
	}
	return nil
}
