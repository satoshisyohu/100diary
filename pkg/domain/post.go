package domain

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/satoshisyohu/pomodoro/pkg/utils/timeutils"
	pb "github.com/satoshisyohu/pomodoro/proto/post"
	"gorm.io/gorm"
	"time"
)

type Post struct {
	PostId    sql.NullString `gorm:"column:post_id;primaryKey"`
	PostDate  sql.NullTime   `gorm:"column:post_date;type:date;primaryKey"`
	UserId    sql.NullString `gorm:"column:user_id"`
	Title     sql.NullString `gorm:"column:title"`
	Content   sql.NullString `gorm:"column:content"`
	FileName  sql.NullString `gorm:"column:file_name"`
	CreatedAt sql.NullTime   `gorm:"column:created_at"`
	UpdatedAt sql.NullTime   `gorm:"column:updated_at"`
}

// NewPost factory method
func NewPost(req *pb.PostCreateRequest, userId string) *Post {
	var fileName *string
	if req.Image != nil && req.FileName != nil {
		f := fmt.Sprintf("%s/%s", timeutils.GenerateStrDate(), req.GetFileName())
		fileName = &f
	}
	p := &Post{
		PostId:   sql.NullString{String: uuid.New().String(), Valid: true},
		PostDate: sql.NullTime{Time: time.Now(), Valid: true},
		UserId:   sql.NullString{String: userId, Valid: true},
		Title:    sql.NullString{String: req.Title, Valid: true},
		Content:  sql.NullString{String: req.Content, Valid: true},
	}
	if fileName != nil {
		p.FileName = sql.NullString{String: *fileName, Valid: true}
	} else {
		p.FileName = sql.NullString{Valid: false}
	}
	return p
}

type IPostService interface {
	Create(tx *gorm.DB, req *pb.PostCreateRequest, userId string) (error, *string)
}
