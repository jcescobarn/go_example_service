package handlers

import (
	"NIST/config"
	"NIST/entities"
	"NIST/repositories"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestUserCreate(t *testing.T) {
	// Configurar la base de datos para la prueba
	enverror := godotenv.Load("../.env")
	if enverror != nil {
		log.Fatal("Error loading .env file")
	}
	config.ConnectToDB()
	config.DB.AutoMigrate(&entities.Users{})
	// Configurar el entorno de prueba
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/users", UserCreate)

	// Crear una solicitud de prueba
	requestBody := UserRequestBody{
		Username: "jcescobarn",
		Name:     "Juan Camilo Escobar",
		Password: "secret_password",
		Email:    "jcescobarn@example.com",
	}
	jsonBody, _ := json.Marshal(requestBody)
	request, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonBody))
	request.Header.Set("Content-Type", "application/json")

	// Crear un contexto de prueba
	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)
	context.Request = request

	// Ejecutar la función
	UserCreate(context)

	// Verificar la respuesta
	response := recorder.Result()
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		t.Errorf("Se esperaba un código de estado %d, pero se obtuvo %d", http.StatusOK, response.StatusCode)
	}

	var responseBody map[string]interface{}
	json.NewDecoder(response.Body).Decode(&responseBody)

	expectedUsername := requestBody.Username
	if responseBody["username"] != expectedUsername {
		t.Errorf("Se esperaba el username '%s', pero se obtuvo '%s'", expectedUsername, responseBody["username"])
	}

	expectedName := requestBody.Name
	if responseBody["name"] != expectedName {
		t.Errorf("Se esperaba el nombre '%s', pero se obtuvo '%s'", expectedName, responseBody["name"])
	}

	expectedEmail := requestBody.Email
	if responseBody["email"] != expectedEmail {
		t.Errorf("Se esperaba el email '%s', pero se obtuvo '%s'", expectedEmail, responseBody["email"])
	}

	expectedStatus := true // Assumiendo que el resultado de CreateUser siempre será true en este caso
	if responseBody["status"] != expectedStatus {
		t.Errorf("Se esperaba el estado '%t', pero se obtuvo '%t'", expectedStatus, responseBody["status"])
	}

	// Borrar los registros de la tabla sin eliminar la tabla
	err := config.DB.Exec("DELETE FROM users").Error
	assert.NoError(t, err)
}

func TestUsersGet(t *testing.T) {
	// Configurar la base de datos para la prueba
	enverror := godotenv.Load("../.env")
	if enverror != nil {
		log.Fatal("Error loading .env file")
	}
	config.ConnectToDB()
	config.DB.AutoMigrate(&entities.Users{})
	// Preparar el entorno de prueba
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/users", UsersGet)

	// Crear usuarios de prueba
	users := []entities.Users{
		{Username: "user1", Name: "User 1", Password: "password1", Email: "user1@example.com"},
		{Username: "user2", Name: "User 2", Password: "password2", Email: "user2@example.com"},
	}

	for _, user := range users {
		_, err := repositories.CreateUser(user.Username, user.Name, user.Password, user.Email)
		if err != nil {
			t.Fatalf("Error al crear usuario de prueba: %v", err)
		}
		defer repositories.DeleteUser(user.Username)
	}

	// Crear una solicitud HTTP GET
	request, _ := http.NewRequest("GET", "/users", nil)

	// Realizar la solicitud HTTP
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Verificar la respuesta
	response := recorder.Result()
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		t.Errorf("Se esperaba un código de estado %d, pero se obtuvo %d", http.StatusOK, response.StatusCode)
	}

	// Leer y verificar el cuerpo de la respuesta
	var responseBody []entities.Users
	err := json.NewDecoder(response.Body).Decode(&responseBody)
	if err != nil {
		t.Errorf("Error al decodificar el cuerpo de la respuesta: %v", err)
	}

	// Verificar la cantidad de usuarios obtenidos
	expectedUsersCount := len(users)
	actualUsersCount := len(responseBody)
	if actualUsersCount != expectedUsersCount {
		t.Errorf("Se esperaba obtener %d usuarios, pero se obtuvieron %d", expectedUsersCount, actualUsersCount)
	}
	// Borrar los registros de la tabla sin eliminar la tabla
	error_db := config.DB.Exec("DELETE FROM users").Error
	assert.NoError(t, error_db)
}

func TestUserGetByUsername(t *testing.T) {
	// Configurar la base de datos para la prueba
	enverror := godotenv.Load("../.env")
	if enverror != nil {
		log.Fatal("Error loading .env file")
	}
	config.ConnectToDB()
	config.DB.AutoMigrate(&entities.Users{})
	// Preparar el entorno de prueba
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/users/:username", UserGetByUsername)

	// Crear un usuario de prueba
	username := "johndoe"
	name := "John Doe"
	password := "password123"
	email := "johndoe@example.com"
	_, err := repositories.CreateUser(username, name, password, email)
	if err != nil {
		t.Fatalf("Error al crear usuario de prueba: %v", err)
	}
	defer repositories.DeleteUser(username)

	// Crear una solicitud HTTP GET
	url := fmt.Sprintf("/users/%s", username)
	request, _ := http.NewRequest("GET", url, nil)

	// Realizar la solicitud HTTP
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Verificar la respuesta
	response := recorder.Result()
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		t.Errorf("Se esperaba un código de estado %d, pero se obtuvo %d", http.StatusOK, response.StatusCode)
	}

	// Leer y verificar el cuerpo de la respuesta
	var responseBody entities.Users
	err = json.NewDecoder(response.Body).Decode(&responseBody)
	if err != nil {
		t.Errorf("Error al decodificar el cuerpo de la respuesta: %v", err)
	}

	// Verificar que los campos del usuario no estén vacíos
	if responseBody.Username != username || responseBody.Name != name || responseBody.Password == "" || responseBody.Email == "" {
		t.Errorf("Se esperaba obtener un usuario con todos los campos no vacíos, pero se obtuvo: %+v", responseBody)
	}
}

func TestUserUpdate(t *testing.T) {
	// Configurar la base de datos para la prueba
	enverror := godotenv.Load("../.env")
	if enverror != nil {
		log.Fatal("Error loading .env file")
	}
	config.ConnectToDB()
	config.DB.AutoMigrate(&entities.Users{})
	// Preparar el entorno de prueba
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.PUT("/users/:username", UserUpdate)

	// Crear un usuario de prueba
	username := "johndoe"
	name := "John Doe"
	password := "password123"
	email := "johndoe@example.com"
	_, err := repositories.CreateUser(username, name, password, email)
	if err != nil {
		t.Fatalf("Error al crear usuario de prueba: %v", err)
	}
	defer repositories.DeleteUser(username)

	// Definir los nuevos datos para actualizar el usuario
	newPassword := "newpassword123"
	newEmail := "newemail@example.com"
	requestBody := UserRequestBody{
		Password: newPassword,
		Email:    newEmail,
	}
	requestBodyBytes, _ := json.Marshal(requestBody)

	// Crear una solicitud HTTP PUT
	url := fmt.Sprintf("/users/%s", username)
	request, _ := http.NewRequest("PUT", url, bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")

	// Realizar la solicitud HTTP
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Verificar la respuesta
	response := recorder.Result()
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		t.Errorf("Se esperaba un código de estado %d, pero se obtuvo %d", http.StatusOK, response.StatusCode)
	}

	// Leer y verificar el cuerpo de la respuesta
	var responseBody gin.H
	err = json.NewDecoder(response.Body).Decode(&responseBody)
	if err != nil {
		t.Errorf("Error al decodificar el cuerpo de la respuesta: %v", err)
	}

	// Verificar los campos de la respuesta
	if updated, ok := responseBody["updated"].(bool); !ok || !updated {
		t.Errorf("Se esperaba que el campo 'updated' en la respuesta fuera verdadero")
	}

	if message, ok := responseBody["message"].(string); !ok || message != "Usuario actualizado con éxito" {
		t.Errorf("Se esperaba que el campo 'message' en la respuesta fuera 'Usuario actualizado con éxito'")
	}

}
