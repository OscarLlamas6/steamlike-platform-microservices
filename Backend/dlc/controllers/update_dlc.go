package controllers

import (
	"dlc-service/configs"
	"dlc-service/models"
	"dlc-service/responses"
	"dlc-service/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateDLC() gin.HandlerFunc {

	return func(c *gin.Context) {

		var DLCAux models.DLCUpdate
		db := configs.ConnectDB()
		defer db.Close()

		//validate the request body
		if err := c.BindJSON(&DLCAux); err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.DLC{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&DLCAux); validationErr != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.DLC{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		// Guardando imagen en bucket s3
		gameImage := DLCAux.ImageURL
		if DLCAux.UpdateImage > 0 {
			newGameImage, awsErr := utils.UploadImage(DLCAux.ImageURL)
			if !awsErr {
				defer db.Close()
				c.JSON(http.StatusBadRequest, responses.DLC{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "Error al guardar portada del DLC"}})
				return
			}
			gameImage = newGameImage
		}

		myQuery, err := db.Prepare("UPDATE `DLC` SET name = ?, imageURL = ?, releaseDate = ?, description = ?, isGlobal = ?, globalPrice = ?, globalDiscount = ? WHERE idDLC = ?")
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.DLC{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		_, err = myQuery.Exec(DLCAux.Name, gameImage, DLCAux.ReleaseDate, DLCAux.Description, DLCAux.IsGlobal, DLCAux.GlobalPrice, DLCAux.GlobalDiscount, DLCAux.IDDLC)
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.DLC{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		// Convirtiendo ID del DLC a string
		currentDLCID := strconv.FormatInt(DLCAux.IDDLC, 10)

		// Actualizar todos los registros maestro-detalle de DLC-Precio-Region
		if len(DLCAux.Prices) > 0 {
			if !utils.UpdatePrices(DLCAux.Prices, currentDLCID) {
				defer db.Close()
				c.JSON(http.StatusBadRequest, responses.DLC{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "Error al guardar precios del DLC"}})
				return
			}
		}

		c.JSON(http.StatusOK, responses.DLC{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"resultado": "DLC actualizado correctamente :D"}})

	}
}
