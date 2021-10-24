package storage

import (
	"context"

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
func CreatePool(connString string) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, err
	}
	config.MinConns = 2
	config.MaxConns = 8
	pool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}
	return pool, nil
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
