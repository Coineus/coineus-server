package model

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type DBUser struct {
	Id           int
	Name         string
	Email        string
	PasswordHash string
}

func (u User) HashPassword() (DBUser, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	return DBUser{
		Id:           u.Id,
		Name:         u.Name,
		Email:        u.Email,
		PasswordHash: string(bytes),
	}, err
}
