package direview

import (
	reviewcontroller "drum-api/internal/adapter/http/controller/review"
	"drum-api/internal/core/domain"
	reviewusecase "drum-api/internal/core/domain/usecase/review"
	reviewrepository "drum-api/internal/infra/gorm/repositories/review"

	"gorm.io/gorm"
)

func DICreateReviewController(conn *gorm.DB) domain.CreateReviewController {
	r := reviewrepository.New(conn)
	u := reviewusecase.New(r)
	c := reviewcontroller.New(u)

	return c
}
