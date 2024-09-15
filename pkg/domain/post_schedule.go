package domain

import "database/sql"

type PostSchedule struct {
	PostScheduleId sql.NullString `gorm:"column:post_schedule_id;primaryKey"`
	DayOfWeek      sql.NullString `gorm:"column:day_of_week"`
	UserId         sql.NullString `gorm:"column:user_id"`
	CreatedAt      sql.NullTime   `gorm:"column:created_at"`
	UpdatedAt      sql.NullTime   `gorm:"column:updated_at"`
}
