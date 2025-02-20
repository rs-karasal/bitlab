package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID             int    `json:"id"`
	Username       string `json:"username"`
	HashedPassword string `json:"hashed_password"`
}

type UserDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.HashedPassword), []byte(password))
	return err == nil
}
