package rule

import (
	"errors"
	"github.com/satoshisyohu/pomodoro/pkg/domain"
	"github.com/satoshisyohu/pomodoro/pkg/domain/repository"
	pb "github.com/satoshisyohu/pomodoro/proto/post"
	"gorm.io/gorm"
)

type PostService struct {
	r repository.PostRepository
}

func NewPostService(r repository.PostRepository) *PostService {
	return &PostService{r: r}
}

// Create if record is not exist, record is registered and return postId
func (ps PostService) Create(tx *gorm.DB, req *pb.PostCreateRequest, userId string) (error, *string) {
	res, err := ps.r.RetrieveByDateAndUserId(tx, userId)
	if err != nil {
		return err, nil
	}
	if res != nil {
		return errors.New("record Already Exist"), nil
	}

	post := domain.NewPost(req, userId)
	if err = ps.r.Save(tx, post); err != nil {
		return err, nil
	}
	return nil, &post.PostId.String
}
