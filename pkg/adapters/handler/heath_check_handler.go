package handler

import (
	"context"
	"github.com/satoshisyohu/pomodoro/pkg/domain/repository"
	"github.com/satoshisyohu/pomodoro/proto/health"
)

// HealthCheckHandler ヘルスチェックのハンドラー
type HealthCheckHandler struct {
	pb.UnimplementedHealthCheckServiceServer
	r repository.HealthCheckRepository
}

// NewHealthCheckHandler  ファクトリ関数
func NewHealthCheckHandler(r repository.HealthCheckRepository) *HealthCheckHandler {
	return &HealthCheckHandler{
		r: r,
	}
}

// HealthCheck サーバのヘルスチェックを行う
func (h *HealthCheckHandler) HealthCheck(_ context.Context, _ *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	if err := h.r.Ping(); err != nil {
		return nil, err
	}
	return &pb.HealthCheckResponse{}, nil
}
