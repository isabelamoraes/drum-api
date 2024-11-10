package drumusecase

import (
	"drum-api/internal/core/domain/applicationerror"
	"drum-api/internal/infra/gorm/model"
)

func (c *createUsecase) Create(drum *model.Drum) error {
	err := c.drumRepo.Create(drum)

	if err != nil {
		return applicationerror.MakeUnprocessableEntityApplicationError(err.Error())
	}

	return nil
}
