package usecases

import (
	"context"
	"errors"
	"github.com/satoshisyohu/pomodoro/pkg/domain"
	pb "github.com/satoshisyohu/pomodoro/proto/post"
	"gorm.io/gorm"
)

type PostCreateInteractor struct {
	db  *gorm.DB
	ps  domain.IPostService
	pts domain.IPostTagService
}

func NewPostCreateInteractor(db *gorm.DB,
	ps domain.IPostService,
	pts domain.IPostTagService) *PostCreateInteractor {
	return &PostCreateInteractor{
		db:  db,
		ps:  ps,
		pts: pts,
	}
}

type IPostCrateInteractor interface {
	Run(ctx context.Context, req *pb.PostCreateRequest) (*pb.PostCreateResponse, error)
}

func (pci *PostCreateInteractor) Run(ctx context.Context, req *pb.PostCreateRequest) (*pb.PostCreateResponse, error) {
	var postId *string
	err := pci.db.Transaction(func(tx *gorm.DB) error {
		err, p := pci.ps.Create(tx, req, "dummy")
		if err != nil {
			return err
		}
		if len(req.Tags) > 5 {
			return errors.New("tag is over 5")
		}
		if len(req.Tags) > 0 {
			// tag分リクエストを作成してレコードを登録する
			if err = pci.pts.Create(tx, req.Tags, p); err != nil {
				return err
			}
		}
		postId = p
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &pb.PostCreateResponse{PostId: *postId}, nil
}
