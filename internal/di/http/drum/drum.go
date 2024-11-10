package didrum

import (
	drumcontroller "drum-api/internal/adapter/http/controller/drum"
	"drum-api/internal/core/domain"
	drumusecase "drum-api/internal/core/domain/usecase/drum"
	drumrepository "drum-api/internal/infra/gorm/repositories/drum"

	"gorm.io/gorm"
)

func DICreateDrumController(conn *gorm.DB) domain.CreateDrumController {
	r := drumrepository.NewCreate(conn)
	u := drumusecase.NewCreate(r)
	c := drumcontroller.NewCreate(u)

	return c
}

func DIListDrumController(conn *gorm.DB) domain.ListDrumController {
	r := drumrepository.NewList(conn)
	u := drumusecase.NewList(r)
	c := drumcontroller.NewList(u)

	return c
}

func DIGetOneByIDDrumController(conn *gorm.DB) domain.GetOneByIDDrumController {
	r := drumrepository.NewGetOneByID(conn)
	u := drumusecase.NewGetOneByID(r)
	c := drumcontroller.NewGetOneByID(u)

	return c
}

func DIDeleteDrumController(conn *gorm.DB) domain.DeleteDrumController {
	r := drumrepository.NewDelete(conn)
	u := drumusecase.NewDelete(r)
	c := drumcontroller.NewDelete(u)

	return c
}

func DIEditDrumController(conn *gorm.DB) domain.EditDrumController {
	r := drumrepository.NewEdit(conn)
	u := drumusecase.NewEdit(r)
	c := drumcontroller.NewEdit(u)

	return c
}
