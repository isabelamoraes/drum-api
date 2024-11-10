package domain

import (
	"drum-api/internal/infra/gorm/model"

	"github.com/gin-gonic/gin"
)

type CreateReviewRepository interface {
	Create(review *model.Review) error
}
type CreateReviewUsecase interface {
	Create(review *model.Review) error
}

type CreateReviewController interface {
	Create(c *gin.Context)
}
