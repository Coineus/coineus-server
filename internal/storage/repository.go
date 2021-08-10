package storage

import "github.com/coineus/coineus-server/model"

type RecentOperationStore interface {
	AddRecentOperation(operation model.RecentOperation) error
	GetAllRecentOperations(userId int) ([]model.RecentOperation, error)
	GetRecentOperationById(userId int, id int) (model.RecentOperation, error)
	DeleteRecentOperation(operation model.RecentOperation) error
	UpdateRecentOperation(operation model.RecentOperation) error
}

type ArchivedOperationStore interface {
	AddArchivedOperation(operation model.ArchivedOperation) error
	GetAllArchivedOperations(userId int) ([]model.ArchivedOperation, error)
	GetArchivedOperationById(userId int, id int) (model.ArchivedOperation, error)
	DeleteArchivedOperation(operation model.ArchivedOperation) error
}

type UserStore interface {
	AddUser(user model.User) error
	GetByMail(mail string) (model.DBUser, error)
	DeleteUser(user model.User) error
	UpdateUser(user model.User) error
}

type WalletStore interface {
	AddWallet(wallet model.Wallet) error
	GetAllWallets(userId int) ([]model.Wallet, error)
	GetWalletById(userId int, id int) (model.Wallet, error)
	DeleteWallet(wallet model.Wallet) error
}

type WalletOperationStore interface {
	AddWalletOperation(walletOperation model.WalletOperation) error
	GetAllWalletOperation(walletId int) ([]model.WalletOperationDTO, error)
	DeleteWalletOperation(walletOperation model.WalletOperation) error
	DeleteAllWalletOperations(walletIdToDelete int) error
}
