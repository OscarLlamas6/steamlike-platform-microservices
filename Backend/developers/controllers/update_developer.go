package controllers

import (
	"net/http"

	"developers-service/configs"
	"developers-service/models"
	"developers-service/responses"
	"developers-service/utils"

	"github.com/gin-gonic/gin"
)

func UpdateDeveloper() gin.HandlerFunc {

	return func(c *gin.Context) {

		var DeveloperAux models.DeveloperComplete
		db := configs.ConnectDB()
		defer db.Close()

		//validate the request body
		if err := c.BindJSON(&DeveloperAux); err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.Developer{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&DeveloperAux); validationErr != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.Developer{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		devImage := DeveloperAux.ImageURL
		if DeveloperAux.UpdateImage > 0 {
			newDevImage, awsErr := utils.UploadImage(DeveloperAux.ImageURL)
			if !awsErr {
				defer db.Close()
				c.JSON(http.StatusBadRequest, responses.Developer{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "Error al guardar portada del juego"}})
				return
			}
			devImage = newDevImage
		}

		myQuery, err := db.Prepare("UPDATE `Developer` SET name = ?, country = ?, imageURL = ? ,email = ? WHERE idDeveloper = ?")
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.Developer{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		_, err = myQuery.Exec(DeveloperAux.Name, DeveloperAux.Pais, devImage, DeveloperAux.Email, DeveloperAux.IdDeveloper)
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.Developer{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.Developer{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"resultado": "Developer actualizado correctamente :D"}})
	}
}
