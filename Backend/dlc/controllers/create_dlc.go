package controllers

import (
	"dlc-service/configs"
	"dlc-service/models"
	"dlc-service/responses"
	"dlc-service/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var (
	validate = validator.New()
)

func CreateDLC() gin.HandlerFunc {

	return func(c *gin.Context) {

		var newDLC models.DLC
		db := configs.ConnectDB()
		defer db.Close()

		//validate the request body
		if err := c.BindJSON(&newDLC); err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.DLC{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&newDLC); validationErr != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.DLC{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		var newDLCImage string = "none"
		if newDLC.ImageURL != "" {
			// Guardando imagen en bucket s3
			gameImage, awsErr := utils.UploadImage(newDLC.ImageURL)
			if !awsErr {
				defer db.Close()
				c.JSON(http.StatusBadRequest, responses.DLC{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "Error al guardar portada del DLC"}})
				return
			}
			newDLCImage = gameImage
		}

		////// REGISTRAR NUEVO DLC
		myQuery, err := db.Prepare("INSERT INTO DLC (name, idGame, isDeleted, imageURL, description, releaseDate, isGlobal, globalPrice, globalDiscount) VALUES(?,?,?,?,?,?,?,?,?)")
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.DLC{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		res, err := myQuery.Exec(newDLC.Name, newDLC.IDGame, 0, newDLCImage, newDLC.Description, newDLC.ReleaseDate, newDLC.IsGlobal, newDLC.GlobalPrice, newDLC.GlobalDiscount)
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.DLC{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		lid, err := res.LastInsertId()
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.DLC{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		// Convirtiendo ID del nuevo DLC a string
		newDLCID := strconv.FormatInt(lid, 10)

		// Guardando todos los registros maestro-detalle de DLC-Precio-Region
		if !utils.SavePrices(newDLC.Prices, newDLCID) {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.DLC{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "Error al guardar precios del DLC"}})
			return
		}

		c.JSON(http.StatusCreated, responses.DLC{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"resultado": "DLC registrado correctamente :D", "id": lid}})

	}
}
