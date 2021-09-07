package recentOperation

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

func (rp *Repository) AddRecentOperation(operation model.RecentOperation) error {
	uuid := app.CreateUUID()
	_, err := rp.db.Exec(context.Background(), "insert into recent_operations (created_at, buy_cost, coin_amount, coin_symbol, user_id, hash_id) values (now(), $1, $2, $3, $4, $5);",
		operation.BuyCost,
		operation.CoinAmount,
		operation.CoinSymbol,
		operation.UserId,
		uuid)
	return err
}

func (rp *Repository) GetAllRecentOperations(userId string) ([]model.RecentOperation, error) {
	var operations []model.RecentOperation

	rows, err := rp.db.Query(context.Background(), "select created_at, buy_cost, coin_amount, coin_symbol, user_id, hash_id from recent_operations where user_id = $1", userId)

	for rows.Next() {
		var operation model.RecentOperation
		err := rows.Scan(
			&operation.CreatedAt,
			&operation.BuyCost,
			&operation.CoinAmount,
			&operation.CoinSymbol,
			&operation.UserId,
			&operation.Id,
		)
		if err != nil {
			return nil, err
		}
		operations = append(operations, operation)
	}
	return operations, err
}

func (rp *Repository) GetRecentOperationById(userId string, id string) (model.RecentOperation, error) {
	var operation model.RecentOperation

	row := rp.db.QueryRow(context.Background(), "select created_at, buy_cost, coin_amount, coin_symbol, user_id, hash_id from recent_operations where hash_id = $1 and user_id = $2", id, userId)
	err := row.Scan(
		&operation.CreatedAt,
		&operation.BuyCost,
		&operation.CoinAmount,
		&operation.CoinSymbol,
		&operation.UserId,
		&operation.Id,
	)
	if err != nil {
		return model.RecentOperation{}, err
	}
	return operation, err
}

func (rp *Repository) DeleteRecentOperation(operation model.RecentOperation) error {
	_, err := rp.db.Exec(context.Background(), "delete from recent_operations where hash_id = $1 and user_id = $2;", operation.Id, operation.UserId)
	return err
}

func (rp *Repository) UpdateRecentOperation(operation model.RecentOperation) error {
	_, err := rp.db.Exec(context.Background(), "update recent_operations set buy_cost = $2, coin_amount = $3, coin_symbol = $4 where hash_id = $1;", operation.Id, operation.BuyCost, operation.CoinAmount, operation.CoinSymbol)
	return err
}
