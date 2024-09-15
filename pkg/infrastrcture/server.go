package infrastrcture

import (
	"fmt"
	"github.com/satoshisyohu/pomodoro/pkg/adapters/handler"
	"github.com/satoshisyohu/pomodoro/pkg/adapters/repository"
	"github.com/satoshisyohu/pomodoro/pkg/domain/rule"
	"github.com/satoshisyohu/pomodoro/pkg/usecases"
	health "github.com/satoshisyohu/pomodoro/proto/health"
	post "github.com/satoshisyohu/pomodoro/proto/post"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func InitServer(s *grpc.Server) {

	//logger := config.NewZeroLogger()
	//
	db, err := repository.InitDb()
	if err != nil {
		fmt.Print(err)
	}
	initHealthCheckServer(s, db)
	initPostCreateServer(s, db)
}

func initHealthCheckServer(s *grpc.Server, db *gorm.DB) {

	health.RegisterHealthCheckServiceServer(s,
		handler.NewHealthCheckHandler(
			repository.NewHeathCheckRepository(db),
		),
	)
}

func initPostCreateServer(s *grpc.Server, db *gorm.DB) {
	post.RegisterPostServiceServer(s,
		handler.NewPostHandler(
			usecases.NewPostCreateInteractor(
				db,
				rule.NewPostService(
					repository.NewPostRepository()),
				rule.NewPostTagService(
					repository.NewTagRepository(),
					repository.NewPostTagRepository(),
				),
			),
		),
	)
}
