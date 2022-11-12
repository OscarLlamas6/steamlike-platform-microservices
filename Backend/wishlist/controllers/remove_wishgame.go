package controllers

import (
	"net/http"

	"wishlist-service/configs"
	"wishlist-service/responses"

	"github.com/gin-gonic/gin"
)

func RemoveWishGame() gin.HandlerFunc {

	return func(c *gin.Context) {

		idMyGame := c.Param("idMyGame")
		db := configs.ConnectDB()
		defer db.Close()

		myQuery, err := db.Prepare("UPDATE `MyGames` SET isWishlist = 0 WHERE idMyGame = ?")
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.MyGame{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		myQuery.Exec(idMyGame)

		c.JSON(http.StatusOK, responses.MyGame{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"resultado": "Juego eliminado de la lista de deseos, esperamos y ya lo hayas comprado :D"}})
	}
}
