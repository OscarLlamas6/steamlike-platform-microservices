package controllers

import (
	"dlc-service/configs"
	"dlc-service/models"
	"dlc-service/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetDLCsByGameNoAuth() gin.HandlerFunc {

	return func(c *gin.Context) {

		regionParam := c.DefaultQuery("region", "none")
		idGame := c.Param("idGame")
		db := configs.ConnectDB()
		defer db.Close()

		//Arreglo para almacenar todos los DLC
		myDLCItemsList := []models.DLCListItem{}

		DLCListQuery, err := db.Query("SELECT idDLC, name, idGame, isDeleted, imageURL, description, releaseDate, isGlobal, globalPrice, globalDiscount FROM DLC WHERE idGame = ? AND isDeleted = 0;", idGame)
		if err != nil {
			defer db.Close()
			successFlag := new(bool)
			*successFlag = false
			c.JSON(http.StatusNotFound, responses.DLCData{Success: successFlag, Data: make([]models.DLCListItem, 0), Message: err.Error()})
			return
		}

		defer DLCListQuery.Close()

		for DLCListQuery.Next() {

			// Pregio y Descuento por region
			var dlcPrice, dlcDiscount float64 = -1, -1

			// Struct para almacenar toda la info del DLC
			myDLCAux := models.DLCListItem{}

			// Arreglo para almacenar los precios por region del DLC
			currentDLCRegionPrices := []models.RegionGameListItem{}

			var id, idGame, isDeleted, isGlobal int64
			var globalPrice, globalDiscount float64
			var name, imageURL, releaseDate, description string

			err = DLCListQuery.Scan(&id, &name, &idGame, &isDeleted, &imageURL, &description, &releaseDate, &isGlobal, &globalPrice, &globalDiscount)
			if err != nil {
				defer db.Close()
				successFlag := new(bool)
				*successFlag = false
				c.JSON(http.StatusNotFound, responses.DLCData{Success: successFlag, Data: make([]models.DLCListItem, 0), Message: err.Error()})
				return
			}

			myDLCAux.IDDLC = id
			myDLCAux.Name = name
			myDLCAux.ImageURL = imageURL
			myDLCAux.ReleaseDate = releaseDate
			myDLCAux.Description = description
			myDLCAux.IDGame = idGame
			myDLCAux.GlobalPrice = globalPrice
			myDLCAux.GlobalDiscount = globalDiscount

			if isDeleted == 0 {
				f := new(bool)
				*f = false

				myDLCAux.IsDeleted = f
			} else {
				t := new(bool)
				*t = true
				myDLCAux.IsDeleted = t
			}

			if isGlobal == 0 {
				f := new(bool)
				*f = false

				myDLCAux.IsGlobal = f
			} else {
				t := new(bool)
				*t = true
				myDLCAux.IsGlobal = t
			}
			// Mapeando Precios por Region del DLC
			pricePerRegionQuery, err := db.Query("SELECT R.`idRegion`, RP.`price`, RP.`discount`, R.`name` FROM RegionPrice AS RP INNER JOIN Region AS R ON RP.idRegion = R.idRegion WHERE RP.idDLC = ? AND RP.isDeleted = 0 AND RP.isDLC = 1;", id)
			if err != nil {
				defer db.Close()
				successFlag := new(bool)
				*successFlag = false
				c.JSON(http.StatusNotFound, responses.DLCData{Success: successFlag, Data: make([]models.DLCListItem, 0), Message: err.Error()})
				return
			}

			defer pricePerRegionQuery.Close()

			for pricePerRegionQuery.Next() {

				var idRegion int64
				var price, discount float64
				var regionName string

				err = pricePerRegionQuery.Scan(&idRegion, &price, &discount, &regionName)
				if err != nil {
					defer db.Close()
					successFlag := new(bool)
					*successFlag = false
					c.JSON(http.StatusNotFound, responses.DLCData{Success: successFlag, Data: make([]models.DLCListItem, 0), Message: err.Error()})
					return
				}

				if regionName == regionParam {
					dlcPrice = price
					dlcDiscount = discount
				}

				regionPriceAux := models.RegionGameListItem{
					IDRegion: idRegion,
					Price:    price,
					Discount: discount,
					Region:   regionName,
				}

				currentDLCRegionPrices = append(currentDLCRegionPrices, regionPriceAux)

			}

			myDLCAux.RegionsAndPrices = currentDLCRegionPrices

			if dlcPrice >= 0 && dlcDiscount >= 0 {
				myDLCAux.GlobalPrice = dlcPrice
				myDLCAux.GlobalDiscount = dlcDiscount
				myDLCItemsList = append(myDLCItemsList, myDLCAux)
			} else if isGlobal > 0 {
				myDLCAux.GlobalPrice = globalPrice
				myDLCAux.GlobalDiscount = globalDiscount
				myDLCItemsList = append(myDLCItemsList, myDLCAux)
			}

		}

		successFlag := new(bool)
		*successFlag = true
		c.JSON(http.StatusOK, responses.DLCData{Success: successFlag, Data: myDLCItemsList, Message: "Info. de DLCs obtenida correctamente :D"})

	}
}
