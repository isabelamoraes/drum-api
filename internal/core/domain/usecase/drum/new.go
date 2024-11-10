package drumusecase

import "drum-api/internal/core/domain"

type createUsecase struct {
	drumRepo domain.CreateDrumRepository
}

func NewCreate(drumRepo domain.CreateDrumRepository) domain.CreateDrumUsecase {
	return &createUsecase{
		drumRepo: drumRepo,
	}
}

type listUsecase struct {
	drumRepo domain.ListDrumRepository
}

func NewList(drumRepo domain.ListDrumRepository) domain.ListDrumUsecase {
	return &listUsecase{
		drumRepo: drumRepo,
	}
}

type getOneByIDUsecase struct {
	drumRepo domain.GetOneByIDDrumRepository
}

func NewGetOneByID(drumRepo domain.GetOneByIDDrumRepository) domain.GetOneByIDDrumUsecase {
	return &getOneByIDUsecase{
		drumRepo: drumRepo,
	}
}

type deleteUsecase struct {
	drumRepo domain.DeleteDrumRepository
}

func NewDelete(drumRepo domain.DeleteDrumRepository) domain.DeleteDrumUsecase {
	return &deleteUsecase{
		drumRepo: drumRepo,
	}
}

type editUsecase struct {
	drumRepo domain.EditDrumRepository
}

func NewEdit(drumRepo domain.EditDrumRepository) domain.EditDrumUsecase {
	return &editUsecase{
		drumRepo: drumRepo,
	}
}
