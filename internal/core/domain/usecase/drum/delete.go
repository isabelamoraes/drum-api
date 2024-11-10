package drumusecase

import (
	"drum-api/internal/core/domain/applicationerror"
)

func (c *deleteUsecase) Delete(ID int) error {
	err := c.drumRepo.Delete(ID)

	if err != nil {
		return applicationerror.MakeUnprocessableEntityApplicationError(err.Error())
	}

	return nil
}
