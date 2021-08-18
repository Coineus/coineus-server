package walletOperation

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

func (rp *Repository) AddWalletOperation(walletOperation model.WalletOperation) error {
	walletOperation.Id = app.CreateUUID()
	_, err := rp.db.Exec(context.Background(), "insert into wallet_operations (operation_id, wallet_id, hash_id) values ($1, $2, $3);", walletOperation.OperationId, walletOperation.WalletId, walletOperation.Id)
	return err
}

func (rp *Repository) GetAllWalletOperation(walletId int) ([]model.WalletOperationDTO, error) {
	var walletOps []model.WalletOperationDTO

	rows, err := rp.db.Query(context.Background(),
		"select wallet_operations.hash_id, wallet_operations.wallet_id, wallet_operations.operation_id, recent_operations.user_id, recent_operations.coin_symbol, recent_operations.coin_amount, recent_operations.buy_cost from wallet_operations inner join recent_operations on wallet_operations.operation_id = recent_operations.hash_id where wallet_operations.wallet_id = $1", walletId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var walletOperation model.WalletOperationDTO

		err := rows.Scan(
			&walletOperation.Id, &walletOperation.WalletId, &walletOperation.OperationId, &walletOperation.UserId, &walletOperation.CoinSymbol, &walletOperation.CoinAmount, &walletOperation.BuyCost)
		if err != nil {
			continue
		}

		walletOps = append(walletOps, walletOperation)
	}
	return walletOps, err

}

func (rp *Repository) DeleteWalletOperation(walletOperation model.WalletOperation) error {
	_, err := rp.db.Exec(context.Background(), "delete from wallet_operations where hash_id = $1 and wallet_id = $2", walletOperation.Id, walletOperation.WalletId)
	return err
}

func (rp *Repository) DeleteAllWalletOperations(walletIdToDelete int) error {
	_, err := rp.db.Exec(context.Background(), "delete from wallet_operations where wallet_id = $1", walletIdToDelete)
	return err
}
