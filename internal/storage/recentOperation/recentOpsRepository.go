package recentOperation

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

func (rp *Repository) AddRecentOperation(operation model.RecentOperation) error {

}

func (rp *Repository) GetAllRecentOperations(userId int) ([]model.RecentOperation, error) {

}

func (rp *Repository) GetRecentOperationById(userId int, id int) (model.RecentOperation, error) {

}

func (rp *Repository) DeleteRecentOperation(operation model.RecentOperation) error {

}

func (rp *Repository) UpdateRecentOperation(operation model.RecentOperation) error {

}
