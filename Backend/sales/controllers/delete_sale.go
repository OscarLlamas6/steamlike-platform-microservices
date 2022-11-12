package controllers

import (
	"net/http"

	"sales-service/configs"
	"sales-service/responses"

	"github.com/gin-gonic/gin"
)

func DeleteSale() gin.HandlerFunc {

	return func(c *gin.Context) {

		idSale := c.Param("idSale")
		db := configs.ConnectDB()
		defer db.Close()

		myQuery, err := db.Prepare("UPDATE `Sale` SET isDeleted = 1 WHERE idSale = ?")
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.Sale{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		myQuery.Exec(idSale)
		c.JSON(http.StatusOK, responses.Sale{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"resultado": "Venta eliminada correctamente :( RIP"}})
	}
}
