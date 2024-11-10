package drumcontroller

import "drum-api/internal/core/domain"

type createDrumController struct {
	usecase domain.CreateDrumUsecase
}

func NewCreate(usecase domain.CreateDrumUsecase) domain.CreateDrumController {
	return &createDrumController{
		usecase: usecase,
	}
}

type listDrumController struct {
	usecase domain.ListDrumUsecase
}

func NewList(usecase domain.ListDrumUsecase) domain.ListDrumController {
	return &listDrumController{
		usecase: usecase,
	}
}

type getOneByIDDrumController struct {
	usecase domain.GetOneByIDDrumUsecase
}

func NewGetOneByID(usecase domain.GetOneByIDDrumUsecase) domain.GetOneByIDDrumController {
	return &getOneByIDDrumController{
		usecase: usecase,
	}
}

type deleteDrumController struct {
	usecase domain.DeleteDrumUsecase
}

func NewDelete(usecase domain.DeleteDrumUsecase) domain.DeleteDrumController {
	return &deleteDrumController{
		usecase: usecase,
	}
}

type editDrumController struct {
	usecase domain.EditDrumUsecase
}

func NewEdit(usecase domain.EditDrumUsecase) domain.EditDrumController {
	return &editDrumController{
		usecase: usecase,
	}
}
