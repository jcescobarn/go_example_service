package repositories_test

import (
	"log"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"NIST/config"
	"NIST/entities"
	"NIST/repositories"
)

func TestCreateUser(t *testing.T) {
	// Configurar la base de datos para la prueba
	enverror := godotenv.Load("../.env")
	if enverror != nil {
		log.Fatal("Error loading .env file")
	}
	config.ConnectToDB()
	config.DB.AutoMigrate(&entities.Users{})

	// Generar datos de prueba
	username := "testuser"
	name := "Test User"
	password := "password123"
	email := "testuser@example.com"

	// Crear el usuario
	result, err := repositories.CreateUser(username, name, password, email)

	// Verificar el resultado de la creación del usuario
	assert.NoError(t, err)
	assert.True(t, result)

	// Borrar los registros de la tabla sin eliminar la tabla
	err = config.DB.Exec("DELETE FROM users").Error
	assert.NoError(t, err)
}

func TestGetUsers(t *testing.T) {
	// Configurar la base de datos para la prueba
	enverror := godotenv.Load("../.env")
	if enverror != nil {
		log.Fatal("Error loading .env file")
	}
	config.ConnectToDB()
	config.DB.AutoMigrate(&entities.Users{})

	// Generar datos de prueba
	users := []entities.Users{
		{
			Model: gorm.Model{
				ID:        7,
				CreatedAt: time.Date(2023, 06, 17, 19, 31, 0, 938212000, time.UTC),
				UpdatedAt: time.Date(2023, 06, 17, 19, 31, 0, 938212000, time.UTC),
				DeletedAt: gorm.DeletedAt{
					Time:  time.Time{},
					Valid: false,
				},
			},
			Username: "jcescobarn",
			Name:     "Juan Camilo Escobar Naranjo",
			Password: "$2a$10$FkKRLIAqpzB2fJ/Z4NHRSOFRSbgJp31w5wQt3Ng7CzTF8jdwZGsbm",
			Email:    "juannaranjo12103@gmail.com",
		},
		{
			Model: gorm.Model{
				ID:        8,
				CreatedAt: time.Date(2023, 06, 17, 19, 32, 0, 938212000, time.UTC),
				UpdatedAt: time.Date(2023, 06, 17, 19, 32, 0, 938212000, time.UTC),
				DeletedAt: gorm.DeletedAt{
					Time:  time.Time{},
					Valid: false,
				},
			},
			Username: "pperez",
			Name:     "pedro perez",
			Password: "$2a$10$FkKRLIAqpzB2fJ/Z4NHRSOFRSbgJp31w5wQt3Ng7CzTF8jdwZGsbm",
			Email:    "pedroperez@example.com",
		},
	}

	// Insertar los usuarios en la base de datos
	for _, user := range users {
		result := config.DB.Create(&user)
		assert.NoError(t, result.Error)
	}

	// Obtener los usuarios
	resultUsers, err := repositories.GetUsers()
	assert.NoError(t, err)

	// Verificar que todos los usuarios estén presentes en los usuarios obtenidos
	for _, user := range users {
		found := false
		for _, u := range resultUsers {
			if u.Username == user.Username && u.Name == user.Name && u.Email == user.Email {
				found = true
				break
			}
		}
		assert.True(t, found)
	}
	// Borrar los registros de la tabla sin eliminar la tabla
	err = config.DB.Exec("DELETE FROM users").Error
	assert.NoError(t, err)

}

func TestGetUserByUsername(t *testing.T) {
	// Configurar la base de datos para la prueba
	enverror := godotenv.Load("../.env")
	if enverror != nil {
		log.Fatal("Error loading .env file")
	}
	config.ConnectToDB()
	config.DB.AutoMigrate(&entities.Users{})

	// Generar datos de prueba
	users := []entities.Users{
		{
			Model: gorm.Model{
				ID:        1,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Username: "johndoe",
			Name:     "John Doe",
			Password: "password123",
			Email:    "johndoe@example.com",
		},
		{
			Model: gorm.Model{
				ID:        2,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Username: "janedoe",
			Name:     "Jane Doe",
			Password: "password456",
			Email:    "janedoe@example.com",
		},
	}

	// Insertar los usuarios en la base de datos
	for _, user := range users {
		result := config.DB.Create(&user)
		assert.NoError(t, result.Error)
	}

	// Caso de prueba: Obtener un usuario existente por su username
	expectedUser := users[0]
	resultUser, err := repositories.GetUserByUsername(expectedUser.Username)
	assert.NoError(t, err)
	assert.Equal(t, expectedUser.Username, resultUser.Username)
	assert.Equal(t, expectedUser.Name, resultUser.Name)
	assert.Equal(t, expectedUser.Email, resultUser.Email)

	// Caso de prueba: Obtener un usuario inexistente por su username
	_, err = repositories.GetUserByUsername("nonexistentuser")
	assert.Error(t, err)
	assert.EqualError(t, err, "no se encontró ningun usuario con el username nonexistentuser")

	// Borrar los registros de la tabla sin eliminar la tabla
	err = config.DB.Exec("DELETE FROM users").Error
	assert.NoError(t, err)
}

func TestUpdateUser(t *testing.T) {
	// Configurar la base de datos para la prueba
	enverror := godotenv.Load("../.env")
	if enverror != nil {
		log.Fatal("Error loading .env file")
	}
	config.ConnectToDB()
	config.DB.AutoMigrate(&entities.Users{})

	// Generar datos de prueba
	username := "johndoe"
	password := "nuevopassword123"
	email := "nuevoemail@example.com"

	// Crear un usuario de prueba
	user := entities.Users{
		Model: gorm.Model{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Username: username,
		Name:     "John Doe",
		Password: "password123",
		Email:    "johndoe@example.com",
	}
	result := config.DB.Create(&user)
	assert.NoError(t, result.Error)

	// Caso de prueba: Actualizar usuario existente
	updateResult, err := repositories.UpdateUser(username, password, email)
	assert.NoError(t, err)
	assert.Equal(t, true, updateResult)

	// Eliminar el usuario de la base de datos
	err = config.DB.Exec("DELETE FROM users").Error
	assert.NoError(t, err)
}

func TestDeleteUser(t *testing.T) {
	// Configurar la base de datos para la prueba
	enverror := godotenv.Load("../.env")
	if enverror != nil {
		log.Fatal("Error loading .env file")
	}
	config.ConnectToDB()
	config.DB.AutoMigrate(&entities.Users{})

	// Crear un usuario de prueba para eliminar
	username := "testuser"
	user := entities.Users{
		Username: username,
	}

	// Insertar el usuario de prueba en la base de datos
	err := config.DB.Create(&user).Error
	assert.NoError(t, err, "Error al crear el usuario de prueba")

	// Ejecutar la función a probar
	success, err := repositories.DeleteUser(username)

	// Verificar los resultados
	assert.NoError(t, err, "Se produjo un error al eliminar el usuario")
	assert.True(t, success, "La eliminación del usuario no fue exitosa")

	// Verificar que el usuario se haya eliminado correctamente de la base de datos
	var deletedUser entities.Users
	err = config.DB.Where("username = ?", username).First(&deletedUser).Error
	assert.Error(t, gorm.ErrRecordNotFound, err, "El usuario no se eliminó correctamente de la base de datos")

}
