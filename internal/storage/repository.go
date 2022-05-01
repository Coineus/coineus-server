package storage

import "github.com/coineus/coineus-server/model"

type RecentOperationStore interface {
	AddRecentOperation(operation model.RecentOperation) error
	GetAllRecentOperations(userId string) ([]model.RecentOperation, error)
	GetRecentOperationById(userId string, id string) (model.RecentOperation, error)
	GetRecentOperationByCoinSymbol(coinSymbol string) (model.RecentOperation, error)
	DeleteRecentOperation(operation model.RecentOperation) error
	UpdateRecentOperation(operation model.RecentOperation) error
}

type ArchivedOperationStore interface {
	AddArchivedOperation(operation model.ArchivedOperation) error
	GetAllArchivedOperations(userId string) ([]model.ArchivedOperation, error)
	GetArchivedOperationById(userId string, id string) (model.ArchivedOperation, error)
	DeleteArchivedOperation(operation model.ArchivedOperation) error
}

type UserStore interface {
	AddUser(user model.User) error
	GetByMail(mail string) (model.UserDTO, error)
	DeleteUser(user model.User) error
	UpdateUser(user model.User) error
}

type WalletStore interface {
	AddWallet(wallet model.Wallet) error
	GetAllWallets(userId string) ([]model.Wallet, error)
	GetWalletById(userId string, id string) (model.Wallet, error)
	DeleteWallet(wallet model.Wallet) error
}

type WalletOperationStore interface {
	AddWalletOperation(walletOperation model.WalletOperation) error
	GetAllWalletOperation(walletId string) ([]model.WalletOperationDTO, error)
	DeleteWalletOperation(walletOperation model.WalletOperation) error
	DeleteAllWalletOperations(walletIdToDelete string) error
}
