package storage

import (
	"context"
	"log"

	"github.com/coineus/coineus-server/internal/storage/archivedOperation"
	"github.com/coineus/coineus-server/internal/storage/recentOperation"
	"github.com/coineus/coineus-server/internal/storage/user"
	"github.com/coineus/coineus-server/internal/storage/wallet"
	"github.com/coineus/coineus-server/internal/storage/walletOperation"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Database struct {
	Db          *pgxpool.Pool
	ArchivedOps ArchivedOperationStore
	RecentOps   RecentOperationStore
	Users       UserStore
	Wallets     WalletStore
	WalletOps   WalletOperationStore
}

//Db connection and configuration
func CreatePool(connString string) *pgxpool.Pool {
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		log.Println("db conn err. parse config fail : ", err)
	}
	config.MinConns = 2
	config.MaxConns = 8
	pool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Println("db conn err : ", err)
	}
	return pool
}

func New(db *pgxpool.Pool) *Database {
	return &Database{
		Db:          db,
		Users:       user.NewRepository(db),
		ArchivedOps: archivedOperation.NewRepository(db),
		RecentOps:   recentOperation.NewRepository(db),
		Wallets:     wallet.NewRepository(db),
		WalletOps:   walletOperation.NewRepository(db),
	}
}
