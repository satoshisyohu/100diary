package rule

import (
	"github.com/satoshisyohu/pomodoro/pkg/domain"
	"github.com/satoshisyohu/pomodoro/pkg/domain/repository"
	pb "github.com/satoshisyohu/pomodoro/proto/post"
	"gorm.io/gorm"
)

type PostTagService struct {
	tr  repository.TagRepository
	ptr repository.PostTagRepository
}

func (pts PostTagService) Create(tx *gorm.DB, req []*pb.Tag, postId *string) error {
	tags := make([]*domain.Tag, 0)
	postTags := make([]*domain.PostTag, 0)
	for _, v := range req {
		t := domain.NewTag(v)
		tags = append(tags, t)
		postTags = append(postTags, domain.NewPostTag(postId, &t.TagId.String))
	}
	if err := pts.tr.Save(tx, tags); err != nil {
		return err
	}
	if err := pts.ptr.Save(tx, postTags); err != nil {
		return err
	}
	return nil

}
