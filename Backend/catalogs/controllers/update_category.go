package controllers

import (
	"net/http"

	"catalogs-service/configs"
	"catalogs-service/models"
	"catalogs-service/responses"

	"github.com/gin-gonic/gin"
)

func UpdateCategory() gin.HandlerFunc {

	return func(c *gin.Context) {

		var GategoryAux models.CategoryUpdate
		db := configs.ConnectDB()
		defer db.Close()

		//validate the request body
		if err := c.BindJSON(&GategoryAux); err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.Catalog{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&GategoryAux); validationErr != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.Catalog{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		myQuery, err := db.Prepare("UPDATE `Category` SET name = ? WHERE idCategory = ?")
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.Catalog{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		_, err = myQuery.Exec(GategoryAux.Name, GategoryAux.IsDeleted, GategoryAux.IDCategory)
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.Catalog{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.Catalog{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"resultado": "Registro actualizado correctamente :D"}})
	}
}
