package drumusecase

import (
	"drum-api/internal/core/domain/applicationerror"
	"drum-api/internal/infra/gorm/model"
)

func (c *editUsecase) Edit(ID int, drum *model.Drum) error {
	err := c.drumRepo.Edit(ID, drum)

	if err != nil {
		return applicationerror.MakeUnprocessableEntityApplicationError(err.Error())
	}

	return nil
}
