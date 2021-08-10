package archivedOperation

import (
	"github.com/coineus/coineus-server/model"
	"github.com/jinzhu/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (rp *Repository) AddArchivedOperation(operation model.ArchivedOperation) error {
	return nil
}

func (rp *Repository) GetAllArchivedOperations(userId int) ([]model.ArchivedOperation, error) {
	return nil, nil

}

func (rp *Repository) GetArchivedOperationById(userId int, id int) (model.ArchivedOperation, error) {
	return model.ArchivedOperation{}, nil

}

func (rp *Repository) DeleteArchivedOperation(operation model.ArchivedOperation) error {
	return nil
}
