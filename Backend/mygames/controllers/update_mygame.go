package controllers

import (
	"net/http"

	"mygames-service/configs"
	"mygames-service/models"
	"mygames-service/responses"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var (
	validate = validator.New()
)

func UpdateMyGame() gin.HandlerFunc {

	return func(c *gin.Context) {

		var MyGameAux models.MyGamesUpdate
		db := configs.ConnectDB()

		//validate the request body
		if err := c.BindJSON(&MyGameAux); err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.MyGame{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&MyGameAux); validationErr != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.MyGame{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		myQuery, err := db.Prepare("UPDATE `MyGame` SET idUser = ?, idGame = ?, isWishlist = ?, isLibray = ? WHERE idMyGame = ?")
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.MyGame{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		_, err = myQuery.Exec(MyGameAux.IDUser, MyGameAux.IDMyGame, MyGameAux.IsWishlist, MyGameAux.IsLibrary, MyGameAux.IDMyGame)
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.MyGame{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		defer db.Close()
		c.JSON(http.StatusOK, responses.MyGame{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"resultado": "Registro actualizado correctamente :D"}})
	}
}
