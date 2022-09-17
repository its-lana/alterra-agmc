package database

import (
	"Task-Code/config"
	"Task-Code/middlewares"
	"Task-Code/models"
	"errors"
)

func GetAllUsers() (interface{}, error) {
	var users []models.User
	if e := config.DB.Find(&users).Error; e != nil {
		return nil, e
	}
	return users, nil
}

func GetUserByID(id string) (interface{}, error) {
	var user models.User

	if e := config.DB.First(&user, id).Error; e != nil {
		return nil, e
	}
	return user, nil
}

func DeleteUserByID(id string) (interface{}, error) {
	var user models.User

	if rowsAff := config.DB.Delete(&user, id).RowsAffected; rowsAff == 0 {
		return nil, errors.New("Failed Delete! User Id not found!")
	}
	return user, nil
}

func LoginUser(user *models.User) (interface{}, error) {
	var err error
	if err = config.DB.Where("email = ? AND password = ? ", user.Email, user.Password).First(user).Error; err != nil {
		return nil, err
	}

	user.Token, err = middlewares.CreateToken(int(user.ID))
	if err != nil {
		return nil, err
	}

	if err := config.DB.Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
