package domain

import (
	"database/sql"
	pb "github.com/satoshisyohu/pomodoro/proto/post"
	"gorm.io/gorm"
)

type PostTag struct {
	PostId    sql.NullString `gorm:"column:post_id;primaryKey"`
	TagId     sql.NullString `gorm:"column:tag_id;primaryKey"`
	CreatedAt sql.NullTime   `gorm:"column:created_at"`
	UpdatedAt sql.NullTime   `gorm:"column:updated_at"`
}

func NewPostTag(postId, tagId *string) *PostTag {
	return &PostTag{
		PostId: sql.NullString{String: *postId, Valid: true},
		TagId:  sql.NullString{String: *tagId, Valid: true},
	}
}

type IPostTagService interface {
	Create(tx *gorm.DB, req []*pb.Tag, postId *string) error
}
