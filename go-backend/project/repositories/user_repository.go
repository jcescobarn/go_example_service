package repositories

import (
	"NIST/config"
	"NIST/entities"
	"NIST/utils"
	"fmt"
)

// Función que crea un registro en la tabla users
func CreateUser(username string, name string, password string, email string) (bool, error) {

	encripted_password, err := utils.HashPassword(password)
	if err != nil {
		return false, err
	}
	user := &entities.Users{Username: username, Name: name, Password: encripted_password, Email: email}
	result := config.DB.Create(&user)

	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

// Función que consulta todos los registros de la tabla users
func GetUsers() ([]entities.Users, error) {
	var users []entities.Users
	result := config.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

// Función que consulta un registro por id en la tabla app
func GetUserByUsername(username string) (entities.Users, error) {
	var user entities.Users
	result := config.DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return user, fmt.Errorf("no se encontró ningun usuario con el username %s", username)
	}
	return user, nil
}

// Función que actualiza un registro por id en la tabla apps
func UpdateUser(username string, password string, email string) (bool, error) {
	encripted_password, err := utils.HashPassword(password)
	if err != nil {
		return false, err
	}
	var user entities.Users
	config.DB.Where("username = ?", username).First(&user)
	data := &entities.Users{Password: encripted_password, Email: email}
	result := config.DB.Model(&user).Updates(data)
	if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}

// Función que elimina un registro por id en la tabla apps
func DeleteUser(username string) (bool, error) {
	var user entities.Users
	result := config.DB.Where("username = ?", username).Delete(&user)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}
