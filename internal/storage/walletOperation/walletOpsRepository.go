package walletOperation

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

func (rp *Repository) AddWalletOperation(walletOperation model.WalletOperation) error {
	return nil
}

func (rp *Repository) GetAllWalletOperation(walletId int) ([]model.WalletOperationDTO, error) {
	return nil, nil

}

func (rp *Repository) DeleteWalletOperation(walletOperation model.WalletOperation) error {
	return nil

}

func (rp *Repository) DeleteAllWalletOperations(walletIdToDelete int) error {
	return nil
}
