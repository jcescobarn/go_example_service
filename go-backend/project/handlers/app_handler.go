package handlers

import (
	"NIST/repositories"
	"net/http"

	"github.com/gin-gonic/gin"

	"strconv"
)

type AppRequestBody struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func AppCreate(c *gin.Context) {
	var body AppRequestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	result, err := repositories.CreateApp(body.Name, body.Description)
	if err != nil {
		c.JSON(500, gin.H{"Error": "Failed to insert", "message": err})
		return
	}
	c.JSON(200, gin.H{"name": body.Name, "description": body.Description, "created": result})
	return
}

func AppGet(c *gin.Context) {
	apps, err := repositories.GetApp()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}
	c.JSON(200, &apps)
	return
}

func AppGetById(c *gin.Context) {
	id, error := strconv.Atoi(c.Param("id"))
	println(error)
	result, err := repositories.GetAppById(int(id))
	if err != nil {
		c.JSON(500, gin.H{"Error": "Hubo un error al consultar la aplicación", "message": err.Error})
		return
	}
	c.JSON(200, &result)
	return
}

func AppUpdate(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(500, gin.H{"Error": "Hubo un error con el id ingresado", "message": err.Error})
		return
	}
	app, err := repositories.GetAppById(int(id))

	body := AppRequestBody{}
	c.BindJSON(&body)
	result, err := repositories.UpdateApp(int(app.ID), body.Name, body.Description)

	if err != nil {
		c.JSON(500, gin.H{"Error": true, "message": "Failed to update"})
		return
	}

	c.JSON(200, gin.H{"updated": result, "message": "Usuario actualizado con éxito"})
	return
}

// Función para eliminar un registro de la tabla apps
func AppDelete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(500, gin.H{"Error": "Hubo un error con el id ingresado", "message": err.Error})
		return
	}
	result, err := repositories.DeleteApp(int(id))
	if err != nil {
		c.JSON(500, gin.H{"Error": "Hubo un error con el id ingresado", "message": err.Error})
		return
	}
	c.JSON(200, gin.H{"deleted": result})
	return
}
