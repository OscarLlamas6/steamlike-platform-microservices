package controllers

import (
	"net/http"

	"mygames-service/configs"
	"mygames-service/responses"

	"github.com/gin-gonic/gin"
)

func DeleteMyGame() gin.HandlerFunc {

	return func(c *gin.Context) {

		idMyGame := c.Param("idMyGame")
		db := configs.ConnectDB()

		myQuery, err := db.Prepare("UPDATE `MyGames` SET isDeleted = 1 WHERE idMyGame = ?")
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.MyGame{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		myQuery.Exec(idMyGame)
		defer db.Close()
		c.JSON(http.StatusOK, responses.MyGame{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"resultado": "Registro eliminado correctamente :( RIP"}})
	}
}
