package controllers

import (
	"net/http"

	"catalogs-service/configs"
	"catalogs-service/models"
	"catalogs-service/responses"

	"github.com/gin-gonic/gin"
)

func UpdateRegion() gin.HandlerFunc {

	return func(c *gin.Context) {

		var RegionAux models.RegionUpdate
		db := configs.ConnectDB()
		defer db.Close()

		//validate the request body
		if err := c.BindJSON(&RegionAux); err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.Catalog{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&RegionAux); validationErr != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.Catalog{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		myQuery, err := db.Prepare("UPDATE `Region` SET name = ? WHERE idRegion = ?")
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.Catalog{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		_, err = myQuery.Exec(RegionAux.Name, RegionAux.IsDeleted, RegionAux.IDRegion)
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.Catalog{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.Catalog{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"resultado": "Registro actualizado correctamente :D"}})
	}
}
