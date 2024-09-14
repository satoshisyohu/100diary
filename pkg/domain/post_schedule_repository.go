package domain

import (
	"github.com/satoshisyohu/pomodoro/pkg/domain/model"
	"gorm.io/gorm"
)

type PostScheduleRepository interface {
	Save(tx *gorm.DB, postSchedule []*model.PostSchedule) error
}
