package controllers

import (
	"net/http"

	"catalogs-service/configs"
	"catalogs-service/responses"

	"github.com/gin-gonic/gin"
)

func DeleteRegion() gin.HandlerFunc {

	return func(c *gin.Context) {

		idRegion := c.Param("idRegion")
		db := configs.ConnectDB()

		myQuery, err := db.Prepare("UPDATE `Region` SET isDeleted = 1 WHERE idRegion = ?")
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.Catalog{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		myQuery.Exec(idRegion)
		defer db.Close()
		c.JSON(http.StatusOK, responses.Catalog{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"resultado": "Registro eliminado correctamente :( RIP"}})
	}
}
