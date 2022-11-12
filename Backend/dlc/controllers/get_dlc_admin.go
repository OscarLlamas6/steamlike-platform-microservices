package controllers

import (
	"dlc-service/configs"
	"dlc-service/models"
	"dlc-service/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetDLCAdmin() gin.HandlerFunc {

	return func(c *gin.Context) {

		idDLC := c.Param("idDLC")
		db := configs.ConnectDB()
		defer db.Close()

		DLCListQuery, err := db.Query("SELECT idDLC, name, idGame, isDeleted, imageURL, description, releaseDate, isGlobal, globalPrice, globalDiscount FROM DLC WHERE idDLC = ?;", idDLC)
		if err != nil {
			defer db.Close()
			successFlag := new(bool)
			*successFlag = false
			c.JSON(http.StatusNotFound, responses.DLCInfo{Success: successFlag, Data: models.DLCListItem{}, Message: err.Error()})
			return
		}

		// Struct para almacenar toda la info del DLC
		myDLCAux := models.DLCListItem{}
		counter := 0

		defer DLCListQuery.Close()

		for DLCListQuery.Next() {

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
				c.JSON(http.StatusNotFound, responses.DLCInfo{Success: successFlag, Data: models.DLCListItem{}, Message: err.Error()})
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
				c.JSON(http.StatusNotFound, responses.DLCInfo{Success: successFlag, Data: models.DLCListItem{}, Message: err.Error()})
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
					c.JSON(http.StatusNotFound, responses.DLCInfo{Success: successFlag, Data: models.DLCListItem{}, Message: err.Error()})
					return
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

			counter++
		}

		if counter <= 0 {
			defer db.Close()
			successFlag := new(bool)
			*successFlag = false
			c.JSON(http.StatusNotFound, responses.DLCInfo{Success: successFlag, Data: myDLCAux, Message: "No existe ningun DLC con ese id"})
			return
		}

		successFlag := new(bool)
		*successFlag = true
		c.JSON(http.StatusOK, responses.DLCInfo{Success: successFlag, Data: myDLCAux, Message: "Info. de DLCs obtenida correctamente :D"})

	}
}
