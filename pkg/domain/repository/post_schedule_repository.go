package repository

import (
	"github.com/satoshisyohu/pomodoro/pkg/domain"
	"gorm.io/gorm"
)

type PostScheduleRepository interface {
	Save(tx *gorm.DB, postSchedule []*domain.PostSchedule) error
}
