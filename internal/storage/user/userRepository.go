package user

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

func (rp *Repository) AddUser(user model.User) error {
	return nil
}

func (rp *Repository) GetByMail(mail string) (model.DBUser, error) {
	return model.DBUser{}, nil
}

func (rp *Repository) DeleteUser(user model.User) error {
	return nil
}

func (rp *Repository) UpdateUser(user model.User) error {
	return nil
}
