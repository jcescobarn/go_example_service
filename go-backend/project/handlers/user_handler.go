package handlers

import (
	"NIST/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserRequestBody struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func UserCreate(c *gin.Context) {
	var body UserRequestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	result, err := repositories.CreateUser(body.Username, body.Name, body.Password, body.Email)
	if err != nil {
		c.JSON(500, gin.H{"Error": "Failed to insert", "message": err})
		return
	}
	c.JSON(200, gin.H{"username": body.Username, "name": body.Name, "email": body.Email, "status": result})
	return
}

func UsersGet(c *gin.Context) {
	users, err := repositories.GetUsers()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}
	c.JSON(200, &users)
	return
}

func UserGetByUsername(c *gin.Context) {
	var username string = c.Param("username")
	result, err := repositories.GetUserByUsername(username)
	if err != nil {
		c.JSON(500, gin.H{"Error": "Hubo un error al consultar el usuario", "message": err.Error()})
		return
	}
	c.JSON(200, &result)
	return
}

func UserUpdate(c *gin.Context) {
	var username string = c.Param("username")

	user, err := repositories.GetUserByUsername(username)

	body := UserRequestBody{}
	c.BindJSON(&body)
	result, err := repositories.UpdateUser(user.Username, body.Password, body.Email)

	if err != nil {
		c.JSON(500, gin.H{"Error": true, "message": "Failed to update"})
		return
	}

	c.JSON(200, gin.H{"updated": result, "message": "Usuario actualizado con éxito"})
	return
}

// Función para eliminar un registro de la tabla apps
func UserDelete(c *gin.Context) {
	var username string = c.Param("username")
	result, err := repositories.DeleteUser(username)
	if err != nil {
		c.JSON(500, gin.H{"Error": "Hubo un error con el id ingresado", "message": err.Error})
		return
	}
	c.JSON(200, gin.H{"deleted": result})
	return
}
