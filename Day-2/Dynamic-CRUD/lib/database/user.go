package database

import (
	"Dynamic-CRUD/config"
	"Dynamic-CRUD/models"
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

	if e := config.DB.Delete(&user, id).Error; e != nil {
		return nil, e
	}
	return user, nil
}
