package repository

import (
	"github.com/satoshisyohu/pomodoro/pkg/domain"
	irepository "github.com/satoshisyohu/pomodoro/pkg/domain/repository"
	"gorm.io/gorm"
)

type tagRepository struct{}

func NewTagRepository() irepository.TagRepository {
	return &tagRepository{}
}

func (a *tagRepository) Save(tx *gorm.DB, tags []*domain.Tag) error {
	if err := tx.Save(tags).Error; err != nil {
		return err
	}
	return nil
}
