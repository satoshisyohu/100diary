package handler

import (
	"context"
	pb "github.com/satoshisyohu/pomodoro/proto/post"
)

// PostHandler handler for post
type PostHandler struct {
	pb.UnimplementedPostServiceServer
}

// NewPostHandler factory function
func NewPostHandler() *PostHandler {
	return &PostHandler{}
}

// Create crate post
func (h *PostHandler) Create(_ context.Context, _ *pb.PostCreateRequest) (*pb.PostCreateResponse, error) {
	// todo usecaseの呼び出し
	return &pb.PostCreateResponse{}, nil
}
