package drumrepository

import (
	"drum-api/internal/adapter/logger"
	"drum-api/internal/error/errorcontext"
	"drum-api/internal/infra/gorm/model"
	"fmt"
)

func (r *drumRepo) Edit(ID int, drum *model.Drum) error {

	err := r.conn.Model(&model.Drum{}).Where("id", ID).UpdateColumns(map[string]interface{}{
		"name": drum.Name, "mark": drum.Mark, "configuration": drum.Configuration, "is_eletronic": drum.IsEletronic,
	}).Error

	if err != nil {
		logger.LogError(logger.ErrorLog{
			Message: fmt.Sprintf(
				"Error trying to edit drum. Error=%v",
				err.Error(),
			),
			Context: errorcontext.DrumRepositoryEdit,
		})

		return err
	}

	return nil
}
