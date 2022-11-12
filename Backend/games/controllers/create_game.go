package controllers

import (
	"games-service/configs"
	"games-service/models"
	"games-service/responses"
	"games-service/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var (
	validate = validator.New()
)

func CreateGame() gin.HandlerFunc {

	return func(c *gin.Context) {

		var newGame models.Game
		db := configs.ConnectDB()
		defer db.Close()

		//validate the request body
		if err := c.BindJSON(&newGame); err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.Game{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&newGame); validationErr != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.Game{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		var newGameImage string = "none"
		if newGame.ImageURL != "" {
			// Guardando imagen en bucket s3
			gameImage, awsErr := utils.UploadImage(newGame.ImageURL)
			if !awsErr {
				defer db.Close()
				c.JSON(http.StatusBadRequest, responses.Game{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "Error al guardar portada del juego"}})
				return
			}
			newGameImage = gameImage
		}

		////// REGISTRAR NUEVO JUEGO

		myQuery, err := db.Prepare("INSERT INTO Game (name, imageURL, releaseDate, restrictionAge, `Game`.`group`, isDeleted, description, isGlobal, globalPrice, globalDiscount) VALUES(?,?,?,?,?,?,?,?,?,?)")
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.Game{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		res, err := myQuery.Exec(newGame.Name, newGameImage, newGame.ReleaseDate, newGame.RestrictionAge, newGame.Group, 0, newGame.Description, newGame.IsGlobal, newGame.GlobalPrice, newGame.GlobalDiscount)
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.Game{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		lid, err := res.LastInsertId()
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.Game{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		// Convirtiendo ID del nuevo juego a string
		newGameID := strconv.FormatInt(lid, 10)

		// Guardando todos los registros maestro-detalle de Juego-Categoria
		if !utils.SaveCategories(newGame.Categories, newGameID) {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.Game{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "Error al guardar categorias del juego"}})
			return
		}

		// Guardando todos los registros maestro-detalle de Juego-Desarrollador
		if !utils.SaveDevelopers(newGame.Developers, newGameID) {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.Game{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "Error al guardar desarrolladores del juego"}})
			return
		}

		// Guardando todos los registros maestro-detalle de Juego-Precio-Region
		if !utils.SavePrices(newGame.Prices, newGameID) {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.Game{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "Error al guardar precios del juego"}})
			return
		}

		c.JSON(http.StatusCreated, responses.Game{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"resultado": "Juego registrado correctamente :D", "id": lid}})

	}
}
