package handler

import (
	"context"
	pb "github.com/satoshisyohu/pomodoro/proto/article"
)

// ArticleHandler handler for article
type ArticleHandler struct {
	pb.UnimplementedArticleServiceServer
}

// NewArticleHandler factory function
func NewArticleHandler() *ArticleHandler {
	return &ArticleHandler{}
}

// Create crate article
func (h *ArticleHandler) Create(_ context.Context, _ *pb.ArticleCreateRequest) (*pb.ArticleCreateResponse, error) {
	// todo usecaseの呼び出し
	return &pb.ArticleCreateResponse{}, nil
}
