package handlers

import (
	"NIST/entities"
	"NIST/repositories"
	"NIST/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type LoginRequestBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var body LoginRequestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	user, err := repositories.GetUserByUsername(body.Username)
	if err != nil {
		c.JSON(500, gin.H{"Error": "Error al consultar el usuario", "message": err})
		return
	}

	compare := utils.ComparePassword(user.Password, body.Password)
	if compare != nil {
		c.JSON(403, gin.H{"Error": "Contraseña incorrecta", "message": err})
		return
	}

	token, err := generateToken(user)
	if err != nil {
		c.JSON(500, gin.H{"Error": "Error al generar el token", "message": err})
		return
	}

	c.JSON(200, gin.H{"Status": "logged", "Token": token})

}

func generateToken(user entities.Users) (string, error) {
	var jwtSecret = []byte("secret_key")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token expira en 24 horas
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
