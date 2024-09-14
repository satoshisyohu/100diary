package repository

import (
	"github.com/satoshisyohu/pomodoro/pkg/domain"
	"github.com/satoshisyohu/pomodoro/pkg/domain/model"
	"gorm.io/gorm"
)

type postScheduleRepository struct{}

func NewPostScheduleRepository() domain.PostScheduleRepository {
	return &postScheduleRepository{}
}

func (a *postScheduleRepository) Save(tx *gorm.DB, postSchedule []*model.PostSchedule) error {
	if err := tx.Save(postSchedule).Error; err != nil {
		return err
	}
	return nil
}
