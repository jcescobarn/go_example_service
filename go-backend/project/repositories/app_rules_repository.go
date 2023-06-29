package repositories

import (
	"NIST/config"
	"NIST/entities"
	"fmt"
)

// Función que crea un registro en la tabla app_rules
func RegisterAppRule(rule_id string, app_id int) (bool, error) {

	app_rule := &entities.AppRules{Rule_id: rule_id, App_id: app_id}
	result := config.DB.Create(&app_rule)

	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

// Función que consulta registros por id de la aplicación en la tabla app_rules
func GetAppRulesById(app_id int) ([]entities.AppRules, error) {
	var app_rules []entities.AppRules
	result := config.DB.Where("app_id = ?", app_id).Find(&app_rules)
	if result.Error != nil {
		return app_rules, fmt.Errorf("no se encontró ninguna aplicación con el id %d", app_id)
	}
	return app_rules, nil
}

// Función que elimina un registro por id en la tabla app_rules
func DeleteAppRule(app_id int, rule_id string) (bool, error) {
	var app_rules entities.AppRules
	result := config.DB.Where("app_id = ? AND rule_id =?", app_id, rule_id).Delete(&app_rules)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func RulesFixed(app_id string) ([]entities.AppRules, error) {
	var app_rules []entities.AppRules
	result := config.DB.Find(&app_rules, app_id)
	if result.Error != nil {
		return app_rules, fmt.Errorf("no se encontró ninguna regla para la aplicación con el id %d", app_id)
	}
	return app_rules, nil
}
