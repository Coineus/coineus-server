package wallet

import (
	"context"

	"github.com/coineus/coineus-server/internal/app"
	"github.com/coineus/coineus-server/model"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func (rp *Repository) AddWallet(wallet model.Wallet) error {
	wallet.Id = app.CreateUUID()
	_, err := rp.db.Exec(context.Background(), "insert into wallets (name, user_id, hash_id) values ($1, $2, $3);",
		wallet.Name,
		wallet.UserId,
		wallet.Id)
	return err
}

func (rp *Repository) GetAllWallets(userId string) ([]model.Wallet, error) {
	var wallets []model.Wallet

	rows, err := rp.db.Query(context.Background(), "select hash_id, name, user_id from wallets where user_id = $1", userId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var wallet model.Wallet

		err := rows.Scan(&wallet.Id, &wallet.Name, &wallet.UserId)
		if err != nil {
			continue
		}
		wallets = append(wallets, wallet)
	}
	return wallets, err

}

func (rp *Repository) GetWalletById(userId string, id string) (model.Wallet, error) {
	var wallet model.Wallet

	row := rp.db.QueryRow(context.Background(), "select hash_id, name, user_id from wallets where user_id = $1 and hash_id = $2", userId, id)
	err := row.Scan(&wallet.Id, &wallet.Name, &wallet.UserId)
	return wallet, err
}

func (rp *Repository) DeleteWallet(wallet model.Wallet) error {
	_, err := rp.db.Exec(context.Background(), "delete from wallets where hash_id = $1 and user_id = $2", wallet.Id, wallet.UserId)
	return err
}
