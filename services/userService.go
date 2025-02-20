package services

import (
	"my_super_project/models"
	"my_super_project/repositories"
)

func RegisterService(user models.UserDTO) error {
	hashedPassword, err := models.HashPassword(user.Password)
	if err != nil {
		return err
	}

	return repositories.SaveUser(user.Username, hashedPassword)
}
