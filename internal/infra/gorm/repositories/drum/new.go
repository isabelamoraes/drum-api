package drumrepository

import (
	"drum-api/internal/core/domain"

	"gorm.io/gorm"
)

type drumRepo struct {
	conn *gorm.DB
}

func NewCreate(conn *gorm.DB) domain.CreateDrumRepository {
	return &drumRepo{
		conn: conn,
	}
}

func NewList(conn *gorm.DB) domain.ListDrumRepository {
	return &drumRepo{
		conn: conn,
	}
}

func NewGetOneByID(conn *gorm.DB) domain.GetOneByIDDrumRepository {
	return &drumRepo{
		conn: conn,
	}
}

func NewDelete(conn *gorm.DB) domain.DeleteDrumRepository {
	return &drumRepo{
		conn: conn,
	}
}

func NewEdit(conn *gorm.DB) domain.EditDrumRepository {
	return &drumRepo{
		conn: conn,
	}
}
