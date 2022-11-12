package controllers

import (
	"net/http"

	"games-service/configs"
	"games-service/responses"

	"github.com/gin-gonic/gin"
)

func RestoreGame() gin.HandlerFunc {

	return func(c *gin.Context) {

		idGame := c.Param("idGame")
		db := configs.ConnectDB()

		myQuery, err := db.Prepare("UPDATE `Game` SET isDeleted = 0 WHERE idGame = ?")
		if err != nil {
			db.Close()
			c.JSON(http.StatusBadRequest, responses.Game{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		myQuery.Exec(idGame)
		db.Close()
		c.JSON(http.StatusOK, responses.Game{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"resultado": "Juego restaurado correctamente :D"}})
	}
}
