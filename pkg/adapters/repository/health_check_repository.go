package repository

import (
	irepository "github.com/satoshisyohu/pomodoro/pkg/domain/repository"
	"gorm.io/gorm"
)

// heathCheckRepository ヘルスチェック用のリポジトリ
type heathCheckRepository struct {
	tx *gorm.DB
}

// NewHeathCheckRepository ファクトリ関数

func NewHeathCheckRepository(tx *gorm.DB) irepository.HealthCheckRepository {
	return &heathCheckRepository{tx: tx}
}

// Ping ピンする
func (h *heathCheckRepository) Ping() error {
	// db情報を取得
	db, err := h.tx.DB()
	if err != nil {
		return err
	}
	// pingする
	if err = db.Ping(); err != nil {
		return err
	}
	return nil
}
