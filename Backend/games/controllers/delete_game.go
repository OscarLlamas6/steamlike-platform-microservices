package controllers

import (
	"net/http"

	"games-service/configs"
	"games-service/responses"

	"github.com/gin-gonic/gin"
)

func DeleteGame() gin.HandlerFunc {

	return func(c *gin.Context) {

		idGame := c.Param("idGame")
		db := configs.ConnectDB()
		defer db.Close()

		myQuery, err := db.Prepare("UPDATE `Game` SET isDeleted = 1 WHERE idGame = ?")
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.Game{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		myQuery.Exec(idGame)
		c.JSON(http.StatusOK, responses.Game{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"resultado": "Juego eliminado correctamente :( RIP"}})
	}
}
