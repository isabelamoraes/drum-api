package drumrepository

import (
	"drum-api/internal/adapter/logger"
	"drum-api/internal/error/errorcontext"
	"drum-api/internal/infra/gorm/model"
	"fmt"
)

func (r *drumRepo) Create(drum *model.Drum) error {
	err := r.conn.Create(&drum).Error

	if err != nil {
		logger.LogError(logger.ErrorLog{
			Message: fmt.Sprintf(
				"Error trying to create drum. Error=%v",
				err.Error(),
			),
			Context: errorcontext.DrumRepositoryCreate,
		})

		return err
	}

	return nil
}
