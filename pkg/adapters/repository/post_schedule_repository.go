package repository

import (
	"github.com/satoshisyohu/pomodoro/pkg/domain"
	irepository "github.com/satoshisyohu/pomodoro/pkg/domain/repository"
	"gorm.io/gorm"
)

type postScheduleRepository struct{}

func NewPostScheduleRepository() irepository.PostScheduleRepository {
	return &postScheduleRepository{}
}

func (a *postScheduleRepository) Save(tx *gorm.DB, postSchedule []*domain.PostSchedule) error {
	if err := tx.Save(postSchedule).Error; err != nil {
		return err
	}
	return nil
}
