package drumrepository

import (
	"drum-api/internal/adapter/logger"
	"drum-api/internal/core/dto"
	"drum-api/internal/error/errorcontext"
	"drum-api/internal/infra/gorm/model"
	"fmt"
)

func (r *drumRepo) List() ([]dto.DrumDTO, error) {
	var drums []model.Drum

	err := r.conn.Find(&drums).Error

	if err != nil {
		logger.LogError(logger.ErrorLog{
			Message: fmt.Sprintf(
				"Error trying to list drum. Error=%v",
				err.Error(),
			),
			Context: errorcontext.DrumRepositoryList,
		})

		return []dto.DrumDTO{}, err
	}

	result := make([]dto.DrumDTO, len(drums))

	for i := range drums {
		result[i] = dto.ToDrumDTO(&drums[i])
	}

	return result, nil
}
