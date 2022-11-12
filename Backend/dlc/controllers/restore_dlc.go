package controllers

import (
	"net/http"

	"dlc-service/configs"
	"dlc-service/responses"

	"github.com/gin-gonic/gin"
)

func RestoreDLC() gin.HandlerFunc {

	return func(c *gin.Context) {

		idDLC := c.Param("idDLC")
		db := configs.ConnectDB()
		defer db.Close()

		myQuery, err := db.Prepare("UPDATE `DLC` SET isDeleted = 0 WHERE idDLC = ?")
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.DLC{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		myQuery.Exec(idDLC)
		c.JSON(http.StatusOK, responses.DLC{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"resultado": "DLC restaurado correctamente :D"}})
	}
}
