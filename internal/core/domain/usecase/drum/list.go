package drumusecase

import (
	"drum-api/internal/core/domain/applicationerror"
	"drum-api/internal/core/dto"
)

func (c *listUsecase) List() ([]dto.DrumDTO, error) {
	drums, err := c.drumRepo.List()

	if err != nil {
		return drums, applicationerror.MakeUnprocessableEntityApplicationError(err.Error())
	}

	return drums, nil
}
