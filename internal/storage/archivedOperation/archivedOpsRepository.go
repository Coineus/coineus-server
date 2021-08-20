package archivedOperation

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

func (rp *Repository) AddArchivedOperation(operation model.ArchivedOperation) error {
	uuid := app.CreateUUID()
	_, err := rp.db.Exec(context.Background(), "insert into archived_operations (archived_at, created_at, sell_cost, buy_cost, coin_amount, coin_symbol, user_id, hash_id) values (now(), $1, $2, $3, $4, $5, $6, $7);",
		operation.CreatedAt.String(),
		operation.SellCost,
		operation.BuyCost,
		operation.CoinAmount,
		operation.CoinSymbol,
		operation.UserId,
		uuid,
	)
	return err
}

func (rp *Repository) GetAllArchivedOperations(userId int) ([]model.ArchivedOperation, error) {
	var operations []model.ArchivedOperation
	rows, err := rp.db.Query(context.Background(), "select archived_at, created_at, sell_cost, buy_cost, coin_amount, coin_symbol, user_id, hash_id from archived_operations where user_id = $1;", userId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var operation model.ArchivedOperation

		err = rows.Scan(&operation.CreatedAt,
			&operation.SellCost,
			&operation.BuyCost,
			&operation.CoinAmount,
			&operation.CoinSymbol,
			&operation.UserId,
			&operation.Id)
		if err != nil {
			return nil, err
		}
		operations = append(operations, operation)
	}
	return operations, err
}

func (rp *Repository) GetArchivedOperationById(userId int, id int) (model.ArchivedOperation, error) {
	var operation model.ArchivedOperation
	row := rp.db.QueryRow(context.Background(), "select archived_at, created_at, sell_cost, buy_cost, coin_amount, coin_symbol, user_id, hash_id from archived_operations where user_id = $1 and hash_id $2;", userId, id)
	err := row.Scan(&operation.CreatedAt,
		&operation.SellCost,
		&operation.BuyCost,
		&operation.CoinAmount,
		&operation.CoinSymbol,
		&operation.UserId,
		&operation.Id)
	if err != nil {
		return model.ArchivedOperation{}, err
	}
	return operation, err
}

func (rp *Repository) DeleteArchivedOperation(operation model.ArchivedOperation) error {
	_, err := rp.db.Exec(context.Background(), "delete from archived_operations where hash_id = $1 and user_id = $2;", operation.Id, operation.UserId)
	return err
}
