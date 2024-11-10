package reviewrepository

import (
	"drum-api/internal/adapter/logger"
	"drum-api/internal/error/errorcontext"
	"drum-api/internal/infra/gorm/model"
	"fmt"
)

func (r *reviewRepo) Create(review *model.Review) error {
	err := r.conn.Create(&review).Error

	if err != nil {
		logger.LogError(logger.ErrorLog{
			Message: fmt.Sprintf(
				"Error trying to create review. Error=%v",
				err.Error(),
			),
			Context: errorcontext.ReviewRepositoryCreate,
		})

		return err
	}

	return nil
}
