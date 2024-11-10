package reviewcontroller

import "drum-api/internal/core/domain"

type reviewController struct {
	usecase domain.CreateReviewUsecase
}

func New(usecase domain.CreateReviewUsecase) domain.CreateReviewController {
	return &reviewController{
		usecase: usecase,
	}
}
