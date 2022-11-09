package impl

import (
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

type RepositoryParams struct {
	DB *gorm.DB
}

func NewRepository(params RepositoryParams) *repository {
	return &repository{
		db: params.DB,
	}
}
