package repository

import (
	"github.com/satoshisyohu/pomodoro/pkg/domain"
	irepository "github.com/satoshisyohu/pomodoro/pkg/domain/repository"
	"gorm.io/gorm"
)

type postTagRepository struct{}

func NewPostTagRepository() irepository.PostTagRepository {
	return &postTagRepository{}
}

func (a *postTagRepository) Save(tx *gorm.DB, postTags []*domain.PostTag) error {
	if err := tx.Save(postTags).Error; err != nil {
		return err
	}
	return nil
}
