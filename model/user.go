package model

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       string `json:"id"`
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type DBUser struct {
	Id           string
	UserName     string
	CreatedAt    time.Time
	Email        string
	PasswordHash string
}

func (u User) HashPassword() (DBUser, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	return DBUser{
		Id:           u.Id,
		UserName:     u.UserName,
		Email:        u.Email,
		PasswordHash: string(bytes),
	}, err
}
