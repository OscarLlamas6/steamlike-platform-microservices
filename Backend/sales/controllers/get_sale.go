package controllers

import (
	"database/sql"
	"net/http"
	"sales-service/configs"
	"sales-service/models"
	"sales-service/responses"

	"github.com/gin-gonic/gin"
)

func GetSale() gin.HandlerFunc {

	return func(c *gin.Context) {

		idSale := c.Param("idSale")
		db := configs.ConnectDB()
		defer db.Close()

		SaleListQuery, err := db.Query("SELECT idSale, idUser, saleDate, total, metododePago FROM Sale WHERE idSale = ? AND isDeleted = 0;", idSale)
		if err != nil {
			defer db.Close()
			successFlag := new(bool)
			*successFlag = false
			c.JSON(http.StatusNotFound, responses.SaleInfo{Success: successFlag, Data: models.SaleListItem{}, Message: err.Error()})
			return
		}

		// Struct para almacenar toda la info de la venta
		mySaleAux := models.SaleListItem{}
		counter := 0

		defer SaleListQuery.Close()

		for SaleListQuery.Next() {

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
				c.JSON(http.StatusNotFound, responses.SaleInfo{Success: successFlag, Data: models.SaleListItem{}, Message: err.Error()})
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
				c.JSON(http.StatusNotFound, responses.SaleInfo{Success: successFlag, Data: models.SaleListItem{}, Message: err.Error()})
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
					c.JSON(http.StatusNotFound, responses.SaleInfo{Success: successFlag, Data: models.SaleListItem{}, Message: err.Error()})
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

			counter++
		}

		if counter <= 0 {
			defer db.Close()
			successFlag := new(bool)
			*successFlag = false
			c.JSON(http.StatusNotFound, responses.SaleInfo{Success: successFlag, Data: models.SaleListItem{}, Message: "No existe ninguna venta con ese id"})
			return
		}

		successFlag := new(bool)
		*successFlag = true
		c.JSON(http.StatusOK, responses.SaleInfo{Success: successFlag, Data: mySaleAux, Message: "Info. de venta obtenida correctamente :D"})

	}
}
