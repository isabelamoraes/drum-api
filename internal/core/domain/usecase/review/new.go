package reviewusecase

import "drum-api/internal/core/domain"

type createUsecase struct {
	reviewRepo domain.CreateReviewRepository
}

func New(reviewRepo domain.CreateReviewRepository) domain.CreateReviewUsecase {
	return &createUsecase{
		reviewRepo: reviewRepo,
	}
}
