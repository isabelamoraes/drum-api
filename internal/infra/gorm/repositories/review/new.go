package reviewrepository

import (
	"drum-api/internal/core/domain"

	"gorm.io/gorm"
)

type reviewRepo struct {
	conn *gorm.DB
}

func New(conn *gorm.DB) domain.CreateReviewRepository {
	return &reviewRepo{
		conn: conn,
	}
}
