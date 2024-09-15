package domain

import (
	"database/sql"
	"github.com/google/uuid"
	pb "github.com/satoshisyohu/pomodoro/proto/post"
)

type Tag struct {
	TagId     sql.NullString `gorm:"column:tag_id;primaryKey"`
	Name      sql.NullString `gorm:"column:name"`
	CreatedAt sql.NullTime   `gorm:"column:created_at"`
	UpdatedAt sql.NullTime   `gorm:"column:updated_at"`
}

func NewTag(tag *pb.Tag) *Tag {
	return &Tag{
		TagId: sql.NullString{
			String: uuid.New().String(),
			Valid:  true,
		},
		Name: sql.NullString{
			String: tag.Tag, Valid: true,
		},
	}
}
