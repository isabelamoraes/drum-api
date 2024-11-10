package drumusecase

import (
	"drum-api/internal/core/domain/applicationerror"
	"drum-api/internal/core/dto"
)

func (c *getOneByIDUsecase) GetOneByID(ID int) (dto.DrumWithReviewDTO, error) {
	drum, err := c.drumRepo.GetOneByID(ID)

	if err != nil {
		return drum, applicationerror.MakeUnprocessableEntityApplicationError(err.Error())
	}

	return drum, nil
}
