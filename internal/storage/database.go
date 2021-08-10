package storage

import (
	"github.com/coineus/coineus-server/internal/config"
	"github.com/coineus/coineus-server/internal/storage/archivedOperation"
	"github.com/coineus/coineus-server/internal/storage/recentOperation"
	"github.com/coineus/coineus-server/internal/storage/user"
	"github.com/coineus/coineus-server/internal/storage/wallet"
	"github.com/coineus/coineus-server/internal/storage/walletOperation"
	"github.com/jinzhu/gorm"
)

type Database struct {
	db          *gorm.DB
	archivedOps ArchivedOperationStore
	recentOps   RecentOperationStore
	users       UserStore
	wallets     WalletStore
	walletOps   WalletOperationStore
}

//Db connection and configuration
func ConnectDB(config *config.DatabaseConfiguration) {

}

func New(db *gorm.DB) *Database {
	return &Database{
		db:          db,
		users:       user.NewRepository(db),
		archivedOps: archivedOperation.NewRepository(db),
		recentOps:   recentOperation.NewRepository(db),
		wallets:     wallet.NewRepository(db),
		walletOps:   walletOperation.NewRepository(db),
	}
}
