package controllers

import (
	"net/http"
	"sales-service/configs"
	"sales-service/models"
	"sales-service/responses"
	"sales-service/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateSale() gin.HandlerFunc {

	return func(c *gin.Context) {

		var SaleAux models.SaleUpdate
		db := configs.ConnectDB()
		defer db.Close()

		//validate the request body
		if err := c.BindJSON(&SaleAux); err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.Sale{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&SaleAux); validationErr != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.Sale{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		myQuery, err := db.Prepare("UPDATE `Sale` SET idUser = ?, saleDate = ?, total = ?, metododePago = ? WHERE idDLC = ?")
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.Sale{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		_, err = myQuery.Exec(SaleAux.IDUser, SaleAux.SaleDate, SaleAux.Total, SaleAux.MetodoDePago, SaleAux.IDSale)
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.Sale{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		// Convirtiendo ID del juego a string
		currentSaleID := strconv.FormatInt(SaleAux.IDSale, 10)

		// Actualizar todos los registros maestro-detalle de Sale-Detalle
		if len(SaleAux.Detalle) > 0 {
			if !utils.UpdateDetails(SaleAux.Detalle, currentSaleID) {
				defer db.Close()
				c.JSON(http.StatusBadRequest, responses.Sale{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "Error al actualizar detalle de la venta"}})
				return
			}
		}

		c.JSON(http.StatusOK, responses.Sale{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"resultado": "Venta actualizada correctamente :D"}})

	}
}
