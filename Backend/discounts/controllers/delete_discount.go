package controllers

import (
	"net/http"

	"discounts-service/configs"
	"discounts-service/responses"

	"github.com/gin-gonic/gin"
)

func DeleteDiscount() gin.HandlerFunc {

	return func(c *gin.Context) {

		idDiscount := c.Param("idDiscount")
		db := configs.ConnectDB()
		defer db.Close()

		myQuery, err := db.Prepare("UPDATE `Gamediscount` SET isDeleted = 1 WHERE idGameDiscount = ?")
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.Discount{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		myQuery.Exec(idDiscount)

		c.JSON(http.StatusOK, responses.Discount{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"resultado": "Descuento eliminado correctamente :( RIP"}})
	}
}
