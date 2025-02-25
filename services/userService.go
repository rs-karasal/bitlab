package services

import (
	"errors"
	"fmt"
	"my_super_project/config"
	"my_super_project/models"
	"my_super_project/repositories"
	"my_super_project/utils/jwt"
)

func RegisterService(user models.UserDTO) error {
	hashedPassword, err := models.HashPassword(user.Password)
	if err != nil {
		return err
	}

	return repositories.SaveUser(user.Username, hashedPassword)
}

func AuthService(user models.UserDTO, cfg *config.Config) (string, error) {
	if user.Username == "" || user.Password == "" {
		return "", errors.New("empty creadentials")
	}

	userDB, err := repositories.GetUserByUsername(user.Username)
	if err != nil {
		return "", nil
	}

	if !userDB.CheckPassword(user.Password) {
		return "", errors.New("error invalid password")
	} else {
		token, err := jwt.GenerateToken(cfg, userDB.Username)
		if err != nil {
			return "", fmt.Errorf("failed to generate token: %w", err)
		}
		return token, nil
	}
}
