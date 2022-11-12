package controllers

import (
	"net/http"
	"wishlist-service/configs"
	"wishlist-service/models"
	"wishlist-service/responses"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var (
	validate = validator.New()
)

func CreateWishlist() gin.HandlerFunc {

	return func(c *gin.Context) {

		var wishlist models.MyGame
		db := configs.ConnectDB()
		defer db.Close()

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

		defer myQuery2.Close()

		if myQuery2.Next() {
			err := myQuery2.Scan(&idUser)
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

		////// AGREGAR JUEGO AL WISHLIST

		myQuery, err := db.Prepare("INSERT INTO MyGames (idUser, idGame, isDeleted, isWishlist, isLibrary) VALUES(?,?,?,?,?)")
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.MyGame{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		res, err := myQuery.Exec(idUser, wishlist.IDGame, 0, 1, 0)
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

		c.JSON(http.StatusCreated, responses.MyGame{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"resultado": "Juego agregado a la lista de deseos correctamente :D", "id": lid}})

	}
}
