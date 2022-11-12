package controllers

import (
	"database/sql"
	"net/http"
	"sales-service/configs"
	"sales-service/models"
	"sales-service/responses"

	"github.com/gin-gonic/gin"
)

func GetSalesByUser() gin.HandlerFunc {

	return func(c *gin.Context) {

		username := c.Param("username")
		db := configs.ConnectDB()
		defer db.Close()

		///////// GETTING USER ID

		myQuery2, err := db.Query("SELECT idUser FROM User WHERE username = ?;", username)
		if err != nil {
			defer db.Close()
			successFlag := new(bool)
			*successFlag = false
			c.JSON(http.StatusNotFound, responses.SaleData{Success: successFlag, Data: make([]models.SaleListItem, 0), Message: err.Error()})
			return
		}

		var currentSaleIdUser int64

		defer myQuery2.Close()

		if myQuery2.Next() {
			err = myQuery2.Scan(&currentSaleIdUser)
			if err != nil {
				defer db.Close()
				successFlag := new(bool)
				*successFlag = false
				c.JSON(http.StatusNotFound, responses.SaleData{Success: successFlag, Data: make([]models.SaleListItem, 0), Message: err.Error()})
				return
			}
		}

		//Arreglo para almacenar todos los detalles
		myDetailItemsList := []models.SaleListItem{}

		SaleListQuery, err := db.Query("SELECT idSale, idUser, saleDate, total, metododePago FROM Sale WHERE idUser = ? AND isDeleted = 0;", currentSaleIdUser)
		if err != nil {
			defer db.Close()
			successFlag := new(bool)
			*successFlag = false
			c.JSON(http.StatusNotFound, responses.SaleData{Success: successFlag, Data: make([]models.SaleListItem, 0), Message: err.Error()})
			return
		}

		defer SaleListQuery.Close()

		for SaleListQuery.Next() {

			// Struct para almacenar toda la info de la venta
			mySaleAux := models.SaleListItem{}

			// Arreglo para almacenar los detalles de la venta
			currentSaleDetails := []models.SaleDetail{}

			var idSale, idUser int64
			var total float64
			var saleDate, metodoDePago string

			err = SaleListQuery.Scan(&idSale, &idUser, &saleDate, &total, &metodoDePago)
			if err != nil {
				defer db.Close()
				successFlag := new(bool)
				*successFlag = false
				c.JSON(http.StatusNotFound, responses.SaleData{Success: successFlag, Data: make([]models.SaleListItem, 0), Message: err.Error()})
				return
			}

			mySaleAux.IDSale = idSale
			mySaleAux.IDUser = idUser
			mySaleAux.SaleDate = saleDate
			mySaleAux.Total = total
			mySaleAux.MetodoDePago = metodoDePago

			// Mapeando Detalles de la venta
			saleDetailsQuery, err := db.Query("SELECT idSaleDetail, idSale, idGame, idDLC, subTotal, isDLC FROM SaleDetail WHERE idSale = ? ;", idSale)
			if err != nil {
				defer db.Close()
				successFlag := new(bool)
				*successFlag = false
				c.JSON(http.StatusNotFound, responses.SaleData{Success: successFlag, Data: make([]models.SaleListItem, 0), Message: err.Error()})
				return
			}

			defer saleDetailsQuery.Close()

			for saleDetailsQuery.Next() {

				var DetailAux models.SaleDetail

				var idSaleDetail, idSale, isDLC int64
				var subTotal float64
				var idGame, idDLC sql.NullInt64

				err = saleDetailsQuery.Scan(&idSaleDetail, &idSale, &idGame, &idDLC, &subTotal, &isDLC)
				if err != nil {
					defer db.Close()
					successFlag := new(bool)
					*successFlag = false
					c.JSON(http.StatusNotFound, responses.SaleData{Success: successFlag, Data: make([]models.SaleListItem, 0), Message: err.Error()})
					return
				}

				DetailAux.IDDetail = idSaleDetail
				DetailAux.IDSale = idSale

				if idGame.Valid {
					DetailAux.IdGame = idGame.Int64
				} else {
					DetailAux.IdGame = 0
				}

				if idDLC.Valid {
					DetailAux.IdDLC = idDLC.Int64
				} else {
					DetailAux.IdDLC = 0
				}

				DetailAux.SubTotal = subTotal
				DetailAux.IsDLC = isDLC

				currentSaleDetails = append(currentSaleDetails, DetailAux)

			}

			mySaleAux.Detalle = currentSaleDetails

			myDetailItemsList = append(myDetailItemsList, mySaleAux)
		}

		successFlag := new(bool)
		*successFlag = true
		c.JSON(http.StatusNotFound, responses.SaleData{Success: successFlag, Data: myDetailItemsList, Message: "Info. de las ventas obtenida correctamente :D"})

	}
}
