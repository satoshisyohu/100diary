package infrastrcture

import (
	"fmt"
	"github.com/satoshisyohu/pomodoro/pkg/adapters/handler"
	"github.com/satoshisyohu/pomodoro/pkg/adapters/repository"
	pb "github.com/satoshisyohu/pomodoro/proto/health"
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
	initHealthCheck(s, db)

}

func initHealthCheck(s *grpc.Server, db *gorm.DB) {

	pb.RegisterHealthCheckServiceServer(s,
		handler.NewHealthCheckHandler(
			repository.NewHeathCheckRepository(db),
		),
	)
}
