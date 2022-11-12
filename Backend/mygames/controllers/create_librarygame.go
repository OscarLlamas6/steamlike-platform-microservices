package controllers

import (
	"mygames-service/configs"
	"mygames-service/models"
	"mygames-service/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateLibraryGame() gin.HandlerFunc {

	return func(c *gin.Context) {

		var wishlist models.MyGame
		db := configs.ConnectDB()

		//validate the request body
		if err := c.BindJSON(&wishlist); err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.MyGame{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&wishlist); validationErr != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.MyGame{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		var idUser int64 = 0

		// OBTENER ID DEL USUARIO SEGUN SU USERNAME

		myQuery2, err := db.Query("SELECT idUser FROM User WHERE username = ?;", wishlist.Username)
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.MyGame{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if !myQuery2.Next() {
			defer db.Close()
			c.JSON(http.StatusNotFound, responses.MyGame{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "Username invalido"}})
			return
		}

		for myQuery2.Next() {
			err = myQuery2.Scan(&idUser)
			if err != nil {
				defer db.Close()
				c.JSON(http.StatusBadRequest, responses.MyGame{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
		}

		////// AGREGAR JUEGO AL WISHLIST

		myQuery, err := db.Prepare("INSERT INTO MyGames (idUser, idGame, isDeleted, isWishlist, isLibrary) VALUES(?,?,?,?,?)")
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.MyGame{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		res, err := myQuery.Exec(idUser, wishlist.IDGame, 0, 0, 1)
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.MyGame{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		lid, err := res.LastInsertId()
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.MyGame{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		defer db.Close()

		c.JSON(http.StatusCreated, responses.MyGame{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"resultado": "Juego agregado a la libreria correctamente :D", "id": lid}})

	}
}
