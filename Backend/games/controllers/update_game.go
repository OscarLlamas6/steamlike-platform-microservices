package controllers

import (
	"games-service/configs"
	"games-service/models"
	"games-service/responses"
	"games-service/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateGame() gin.HandlerFunc {

	return func(c *gin.Context) {

		var gameAux models.GamesUpdate
		db := configs.ConnectDB()
		defer db.Close()

		//validate the request body
		if err := c.BindJSON(&gameAux); err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.Game{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&gameAux); validationErr != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.Game{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		// Guardando imagen en bucket s3

		gameImage := gameAux.ImageURL
		if gameAux.UpdateImage > 0 {
			newGameImage, awsErr := utils.UploadImage(gameAux.ImageURL)
			if !awsErr {
				defer db.Close()
				c.JSON(http.StatusBadRequest, responses.Game{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "Error al guardar portada del juego"}})
				return
			}
			gameImage = newGameImage
		}

		myQuery, err := db.Prepare("UPDATE `Game` SET name = ?, imageURL = ?, releaseDate = ?, restrictionAge = ?, `Game`.`group` = ?, description = ?, isGlobal = ?, globalPrice = ?, globalDiscount = ? WHERE idGame = ?")
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.Game{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		_, err = myQuery.Exec(gameAux.Name, gameImage, gameAux.ReleaseDate, gameAux.RestrictionAge, gameAux.Group, gameAux.Description, gameAux.IsGlobal, gameAux.GlobalPrice, gameAux.GlobalDiscount, gameAux.IDGame)
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.Game{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		// Convirtiendo ID del juego a string
		actualGameID := strconv.FormatInt(gameAux.IDGame, 10)

		// Actualizar todos los registros maestro-detalle de Juego-Categoria
		if len(gameAux.Categories) > 0 {
			if !utils.UpdateCategories(gameAux.Categories, actualGameID) {
				defer db.Close()
				c.JSON(http.StatusBadRequest, responses.Game{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "Error al guardar categorias del juego"}})
				return
			}
		}

		// Actualizar todos los registros maestro-detalle de Juego-Desarrollador
		if len(gameAux.Developers) > 0 {
			if !utils.UpdateDevelopers(gameAux.Developers, actualGameID) {
				defer db.Close()
				c.JSON(http.StatusBadRequest, responses.Game{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "Error al guardar desarrolladores del juego"}})
				return
			}
		}

		// Actualizar todos los registros maestro-detalle de Juego-Precio-Region
		if len(gameAux.Prices) > 0 {
			if !utils.UpdatePrices(gameAux.Prices, actualGameID) {
				defer db.Close()
				c.JSON(http.StatusBadRequest, responses.Game{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "Error al guardar precios del juego"}})
				return
			}
		}

		c.JSON(http.StatusOK, responses.Game{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"resultado": "Juego actualizado correctamente :D"}})

	}
}
