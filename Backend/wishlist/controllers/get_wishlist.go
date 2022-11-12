package controllers

import (
	"encoding/json"
	"net/http"

	"wishlist-service/configs"
	"wishlist-service/models"
	"wishlist-service/responses"

	"github.com/gin-gonic/gin"
)

func GetWishlist() gin.HandlerFunc {

	return func(c *gin.Context) {

		username := c.Param("username")
		db := configs.ConnectDB()
		defer db.Close()

		var idUsername int64 = 0

		// OBTENER ID DEL USUARIO SEGUN SU USERNAME

		myQuery2, err := db.Query("SELECT idUser FROM User WHERE username = ?;", username)
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.MyGame{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		defer myQuery2.Close()

		if myQuery2.Next() {
			err := myQuery2.Scan(&idUsername)
			if err != nil {
				defer db.Close()
				c.JSON(http.StatusBadRequest, responses.MyGame{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
		} else {
			defer db.Close()
			c.JSON(http.StatusNotFound, responses.MyGame{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "Username invalido"}})
			return
		}

		myQuery, err := db.Query("SELECT * FROM MyGames WHERE idUser = ? AND isWishlist = 1 AND isLibrary = 0 AND isDeleted = 0", idUsername)
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.MyGame{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		defer myQuery.Close()

		var MyGame models.MyGamesUpdate
		myWishlist := []models.MyGamesUpdate{}

		for myQuery.Next() {
			var id, idUser, idGame, isDeleted, isWishlist, isLibrary int64

			err = myQuery.Scan(&id, &idUser, &idGame, &isDeleted, &isWishlist, &isLibrary)
			if err != nil {
				defer db.Close()
				c.JSON(http.StatusBadRequest, responses.MyGame{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}

			MyGame.IDMyGame = id
			MyGame.IDUser = idUser
			MyGame.IDGame = idGame
			MyGame.IsWishlist = isWishlist
			MyGame.IsLibrary = isLibrary

			if isDeleted == 0 {
				f := new(bool)
				*f = false

				MyGame.IsDeleted = f
			} else {
				t := new(bool)
				*t = true
				MyGame.IsDeleted = t
			}

			myWishlist = append(myWishlist, MyGame)
		}

		var myFullWishlist []map[string]interface{}
		wishListJson, err := json.Marshal(myWishlist)
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.MyGame{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		json.Unmarshal(wishListJson, &myFullWishlist)

		c.JSON(http.StatusOK, responses.MyGames{Status: http.StatusOK, Message: "success", Data: myFullWishlist})

	}
}
