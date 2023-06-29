package utils

import (
	"NIST/entities"
	"math"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func ComparePassword(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	}
	return nil
}

func CalculateTotalPages(totalResults int, resultsPerPage int) int {
	return int(math.Ceil(float64(totalResults) / float64(resultsPerPage)))
}

func GetDifference(arr1 []entities.AppRules, arr2 []RuleResponse) []RuleResponse {
	// Crear un mapa para almacenar los IDs del primer arreglo
	ids := make(map[string]struct{})

	// Almacenar los IDs del primer arreglo en el mapa
	for _, obj := range arr1 {
		ids[obj.Rule_id] = struct{}{}
	}

	// Filtrar los objetos del segundo arreglo cuyos IDs no están en el mapa
	var difference []RuleResponse
	for _, obj := range arr2 {
		if _, found := ids[obj.Id]; !found {
			difference = append(difference, obj)
		}
	}

	return difference
}
