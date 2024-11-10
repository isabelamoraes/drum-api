package domain

import (
	"drum-api/internal/core/dto"
	"drum-api/internal/infra/gorm/model"

	"github.com/gin-gonic/gin"
)

type CreateDrumRepository interface {
	Create(drum *model.Drum) error
}

type CreateDrumUsecase interface {
	Create(drum *model.Drum) error
}

type CreateDrumController interface {
	Create(c *gin.Context)
}

type ListDrumRepository interface {
	List() ([]dto.DrumDTO, error)
}

type ListDrumUsecase interface {
	List() ([]dto.DrumDTO, error)
}

type ListDrumController interface {
	List(c *gin.Context)
}

type GetOneByIDDrumRepository interface {
	GetOneByID(ID int) (dto.DrumWithReviewDTO, error)
}

type GetOneByIDDrumUsecase interface {
	GetOneByID(ID int) (dto.DrumWithReviewDTO, error)
}

type GetOneByIDDrumController interface {
	GetOneByID(c *gin.Context)
}

type DeleteDrumRepository interface {
	Delete(ID int) error
}

type DeleteDrumUsecase interface {
	Delete(ID int) error
}

type DeleteDrumController interface {
	Delete(c *gin.Context)
}

type EditDrumRepository interface {
	Edit(ID int, drum *model.Drum) error
}

type EditDrumUsecase interface {
	Edit(ID int, drum *model.Drum) error
}

type EditDrumController interface {
	Edit(c *gin.Context)
}
