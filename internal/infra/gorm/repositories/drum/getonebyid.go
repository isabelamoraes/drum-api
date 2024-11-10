package drumrepository

import (
	"drum-api/internal/adapter/logger"
	"drum-api/internal/core/dto"
	"drum-api/internal/error/errorcontext"
	"drum-api/internal/infra/gorm/model"
	"fmt"
)

func (r *drumRepo) GetOneByID(ID int) (dto.DrumWithReviewDTO, error) {
	var drum model.Drum

	err := r.conn.Preload("Reviews").First(&drum, ID).Error

	if err != nil {
		logger.LogError(logger.ErrorLog{
			Message: fmt.Sprintf(
				"Error trying to get drum. Error=%v",
				err.Error(),
			),
			Context: errorcontext.DrumRepositoryGetOneByID,
		})

		return dto.DrumWithReviewDTO{}, err
	}

	result := dto.ToDrumWithReviewDTO(&drum)

	return result, nil
}
