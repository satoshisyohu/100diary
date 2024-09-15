package repository

import (
	"errors"
	"github.com/satoshisyohu/pomodoro/pkg/domain"
	irepository "github.com/satoshisyohu/pomodoro/pkg/domain/repository"
	"gorm.io/gorm"
)

type PostRepository struct{}

func NewPostRepository() irepository.PostRepository {
	return &PostRepository{}
}

func (a *PostRepository) Save(tx *gorm.DB, Post *domain.Post) error {
	if err := tx.Save(Post).Error; err != nil {
		return err
	}
	return nil
}

func (a *PostRepository) RetrieveByDateAndUserId(tx *gorm.DB, userId string) (*domain.Post, error) {
	var post *domain.Post
	if res := tx.Where("(post_date = current_date() and user_id = ?)", userId).First(&post); res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, res.Error
	}
	return post, nil
}
