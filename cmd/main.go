package main

import (
	"github.com/joho/godotenv"
	"github.com/satoshisyohu/pomodoro/pkg/infrastrcture"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const port = ":50052"

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err = godotenv.Load("../.env"); err != nil {
		log.Fatalf(err.Error())
	}

	s := grpc.NewServer(
	//grpc.UnaryInterceptor(idTokenInterceptor),
	)
	infrastrcture.InitServer(s)
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
