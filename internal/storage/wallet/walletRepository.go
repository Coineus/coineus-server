package wallet

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

func (rp *Repository) AddWallet(wallet model.Wallet) error {
	return nil
}

func (rp *Repository) GetAllWallets(userId int) ([]model.Wallet, error) {
	return nil, nil

}

func (rp *Repository) GetWalletById(userId int, id int) (model.Wallet, error) {
	return model.Wallet{}, nil
}

func (rp *Repository) DeleteWallet(wallet model.Wallet) error {
	return nil
}
