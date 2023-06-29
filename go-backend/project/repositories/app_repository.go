package repositories

import (
	"NIST/config"
	"NIST/entities"
	"fmt"
)

// Función que crea un registro en la tabla apps
func CreateApp(name string, description string) (bool, error) {

	app := &entities.Apps{Name: name, Description: description}
	result := config.DB.Create(&app)

	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

// Función que consulta todos los registros de la tabla apps
func GetApp() ([]entities.Apps, error) {
	var apps []entities.Apps
	result := config.DB.Find(&apps)
	if result.Error != nil {
		return nil, result.Error
	}
	return apps, nil
}

// Función que consulta un registro por id en la tabla app
func GetAppById(id int) (entities.Apps, error) {
	var app entities.Apps
	result := config.DB.First(&app, id)
	if result.Error != nil {
		return app, fmt.Errorf("no se encontró ninguna aplicación con el id %d", id)
	}
	return app, nil
}

// Función que actualiza un registro por id en la tabla apps
func UpdateApp(id int, name string, description string) (bool, error) {

	var app entities.Apps
	config.DB.First(&app, id)
	data := &entities.Apps{Name: name, Description: description}
	result := config.DB.Model(&app).Updates(data)
	if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}

// Función que elimina un registro por id en la tabla apps
func DeleteApp(id int) (bool, error) {
	var app entities.Apps
	result := config.DB.Delete(&app, id)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}
