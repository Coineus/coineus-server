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

type UserDTO struct {
	Id           string    `json:"id"`
	UserName     string    `json:"username"`
	CreatedAt    time.Time `json:"created_at"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password"`
}

func (u User) HashPassword() (UserDTO, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	return UserDTO{
		Id:           u.Id,
		UserName:     u.UserName,
		Email:        u.Email,
		PasswordHash: string(bytes),
	}, err
}
