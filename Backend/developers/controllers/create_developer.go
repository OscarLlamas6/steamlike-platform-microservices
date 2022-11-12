package controllers

import (
	"developers-service/configs"
	"developers-service/models"
	"developers-service/responses"
	"developers-service/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var (
	validate = validator.New()
)

func CreateDeveloper() gin.HandlerFunc {

	return func(c *gin.Context) {

		var newDeveloper models.Developer
		db := configs.ConnectDB()

		//validate the request body
		if err := c.BindJSON(&newDeveloper); err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.Developer{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&newDeveloper); validationErr != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.Developer{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		var newDevImage string = "none"
		if newDeveloper.ImageURL != "" {
			// Guardando imagen en bucket s3
			devImage, awsErr := utils.UploadImage(newDeveloper.ImageURL)
			if !awsErr {
				defer db.Close()
				c.JSON(http.StatusBadRequest, responses.Developer{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "Error al guardar portada del juego"}})
				return
			}
			newDevImage = devImage
		}

		////// REGISTRAR NUEVO DESARROLLADOR

		myQuery, err := db.Prepare("INSERT INTO Developer (name, country, imageURL, email, isDeleted) VALUES(?,?,?,?,?)")
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.Developer{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		res, err := myQuery.Exec(newDeveloper.Name, newDeveloper.Pais, newDevImage, newDeveloper.Email, 0)
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.Developer{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		lid, err := res.LastInsertId()
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.Developer{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		defer db.Close()

		c.JSON(http.StatusCreated, responses.Developer{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"resultado": "Desarrollador registrado correctamente :D", "id": lid}})

	}
}
