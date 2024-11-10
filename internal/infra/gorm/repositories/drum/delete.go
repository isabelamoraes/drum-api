package drumrepository

import (
	"drum-api/internal/adapter/logger"
	"drum-api/internal/error/errorcontext"
	"drum-api/internal/infra/gorm/model"
	"fmt"
)

func (r *drumRepo) Delete(ID int) error {
	var drum model.Drum

	err := r.conn.Delete(&drum, ID).Error

	if err != nil {
		logger.LogError(logger.ErrorLog{
			Message: fmt.Sprintf(
				"Error trying to delete drum. Error=%v",
				err.Error(),
			),
			Context: errorcontext.DrumRepositoryDelete,
		})

		return err
	}

	return nil
}
