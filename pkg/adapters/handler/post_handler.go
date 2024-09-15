package handler

import (
	"context"
	"github.com/satoshisyohu/pomodoro/pkg/usecases"
	pb "github.com/satoshisyohu/pomodoro/proto/post"
)

// PostHandler handler for post
type PostHandler struct {
	pb.UnimplementedPostServiceServer
	u usecases.IPostCrateInteractor
}

// NewPostHandler factory function
func NewPostHandler(u usecases.IPostCrateInteractor) *PostHandler {
	return &PostHandler{u: u}
}

// Create crate post
func (p *PostHandler) Create(ctx context.Context, req *pb.PostCreateRequest) (*pb.PostCreateResponse, error) {
	return p.u.Run(ctx, req)
}
