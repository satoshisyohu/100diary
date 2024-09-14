package repository

import (
	"github.com/satoshisyohu/pomodoro/pkg/domain"
	"github.com/satoshisyohu/pomodoro/pkg/domain/model"
	"gorm.io/gorm"
)

type tagRepository struct{}

func NewTagRepository() domain.TagRepository {
	return &tagRepository{}
}

func (a *tagRepository) Save(tx *gorm.DB, tags []*model.Tag) error {
	if err := tx.Save(tags).Error; err != nil {
		return err
	}
	return nil
}
