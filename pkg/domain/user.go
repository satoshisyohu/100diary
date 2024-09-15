package domain

import "database/sql"

type User struct {
	UserId       sql.NullString `gorm:"column:user_id;primaryKey"`
	LastPostedAt sql.NullString `gorm:"column:last_posted_at"`
	DarkMode     sql.NullTime   `gorm:"column:dark_mode"`
	CreatedAt    sql.NullTime   `gorm:"column:created_at"`
	UpdatedAt    sql.NullTime   `gorm:"column:updated_at"`
}
