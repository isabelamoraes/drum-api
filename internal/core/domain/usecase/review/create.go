package reviewusecase

import (
	"drum-api/internal/core/domain/applicationerror"
	"drum-api/internal/infra/gorm/model"
)

func (c *createUsecase) Create(review *model.Review) error {
	err := c.reviewRepo.Create(review)

	if err != nil {
		return applicationerror.MakeUnprocessableEntityApplicationError(err.Error())
	}

	return nil
}
