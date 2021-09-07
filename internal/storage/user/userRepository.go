package user

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

func (rp *Repository) AddUser(user model.User) error {
	var dbUser model.DBUser
	dbUser, err := user.HashPassword()
	dbUser.Id = app.CreateUUID()
	if err != nil {
		return err
	}

	_, err = rp.db.Exec(context.Background(), "insert into users (password_hash, email, created_at, username, hash_id) values ($1, $2, now(), $3, $4);",
		dbUser.PasswordHash,
		dbUser.Email,
		dbUser.UserName,
		dbUser.Id,
	)
	return err
}

func (rp *Repository) GetByMail(mail string) (model.DBUser, error) {
	var dbUser model.DBUser
	row := rp.db.QueryRow(context.Background(), "select password_hash, email, created_at, username, hash_id from users where email = $1 ;", mail)
	err := row.Scan(&dbUser.PasswordHash, &dbUser.Email, &dbUser.CreatedAt, &dbUser.UserName, &dbUser.Id)
	return dbUser, err
}

func (rp *Repository) DeleteUser(user model.User) error {
	_, err := rp.db.Exec(context.Background(), "delete from users where hash_id = $1;", user.Id)
	return err
}

func (rp *Repository) UpdateUser(user model.User) error {
	dbUser, _ := user.HashPassword()
	_, err := rp.db.Exec(context.Background(), "update users set username = $2, password_hash = $3 where hash_id = $1;", dbUser.Id, dbUser.UserName, dbUser.PasswordHash)
	return err
}
