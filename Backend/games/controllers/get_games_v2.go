package controllers

import (
	"fmt"
	"games-service/configs"
	"games-service/models"
	"games-service/responses"
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func GetGamesv2() gin.HandlerFunc {

	return func(c *gin.Context) {

		claims := jwt.ExtractClaims(c)
		regionParam := claims["region"].(string)
		fmt.Println(regionParam)
		db := configs.ConnectDB()
		defer db.Close()

		//Arreglo para almacenar todos los juegos
		myGamesItemsList := []models.GameListItem{}

		gameListQuery, err := db.Query("SELECT * FROM Game WHERE isDeleted = 0;")
		if err != nil {
			defer db.Close()
			successFlag := new(bool)
			*successFlag = false
			c.JSON(http.StatusNotFound, responses.GamesList{Success: successFlag, Data: make([]models.GameListItem, 0), Message: err.Error()})
			return
		}

		defer gameListQuery.Close()

		for gameListQuery.Next() {

			// Pregio y Descuento por region
			var gamePrice, gameDiscount float64 = -1, -1

			// Struct para almacenar toda la info del juego
			myGameAux := models.GameListItem{}

			// Arreglo para almacenar las categorias del juego
			currentGameCategories := []string{}

			// Arreglo para almacenar las regiones del juego
			currentGameRegions := []string{}

			// Arreglo para almacenar los precios por region del juego
			currentGameRegionPrices := []models.RegionGameListItem{}

			// Arrelo para almacenar los desarrolladores del juego
			currentGameDevelopers := []models.DeveloperGameListItem{}

			var id, isDeleted, group, isGlobal int64
			var globalPrice, globalDiscount float64
			var name, imageURL, releaseDate, restrictionAge, description string

			err = gameListQuery.Scan(&id, &name, &imageURL, &releaseDate, &restrictionAge, &group, &isDeleted, &description, &isGlobal, &globalPrice, &globalDiscount)
			if err != nil {
				defer db.Close()
				successFlag := new(bool)
				*successFlag = false
				c.JSON(http.StatusNotFound, responses.GamesList{Success: successFlag, Data: make([]models.GameListItem, 0), Message: err.Error()})
				return
			}

			myGameAux.GameID = id
			myGameAux.Name = name
			myGameAux.Image = imageURL
			myGameAux.ReleaseDate = releaseDate
			myGameAux.RestrictionAge = restrictionAge
			myGameAux.Description = description
			myGameAux.Group = group

			if isDeleted == 0 {
				f := new(bool)
				*f = false

				myGameAux.IsDeleted = f
			} else {
				t := new(bool)
				*t = true
				myGameAux.IsDeleted = t
			}

			if isGlobal == 0 {
				f := new(bool)
				*f = false

				myGameAux.IsGlobal = f
			} else {
				t := new(bool)
				*t = true
				myGameAux.IsGlobal = t
			}

			// Mapeando Categorias del juego
			gameCategoryQuery, err := db.Query("SELECT `Category`.`name` FROM GameCategory INNER JOIN Category ON GameCategory.idCategory = Category.idCategory WHERE `GameCategory`.`idGame` = ? AND `GameCategory`.isDeleted = 0;", id)
			if err != nil {
				defer db.Close()
				successFlag := new(bool)
				*successFlag = false
				c.JSON(http.StatusNotFound, responses.GamesList{Success: successFlag, Data: make([]models.GameListItem, 0), Message: err.Error()})
				return
			}

			defer gameCategoryQuery.Close()

			for gameCategoryQuery.Next() {

				var categoryName string

				err = gameCategoryQuery.Scan(&categoryName)
				if err != nil {
					defer db.Close()
					successFlag := new(bool)
					*successFlag = false
					c.JSON(http.StatusNotFound, responses.GamesList{Success: successFlag, Data: make([]models.GameListItem, 0), Message: err.Error()})
					return
				}

				currentGameCategories = append(currentGameCategories, categoryName)

			}

			myGameAux.Category = currentGameCategories

			// Mapeando Precios por Region del juego
			pricePerRegionQuery, err := db.Query("SELECT RP.`price`, RP.`discount`, R.`name` FROM RegionPrice AS RP INNER JOIN Region AS R ON RP.idRegion = R.idRegion WHERE RP.idGame = ? AND RP.isDeleted = 0 AND RP.isDLC = 0;", id)
			if err != nil {
				defer db.Close()
				successFlag := new(bool)
				*successFlag = false
				c.JSON(http.StatusNotFound, responses.GamesList{Success: successFlag, Data: make([]models.GameListItem, 0), Message: err.Error()})
				return
			}

			defer pricePerRegionQuery.Close()

			for pricePerRegionQuery.Next() {

				var price, discount float64
				var regionName string

				err = pricePerRegionQuery.Scan(&price, &discount, &regionName)
				if err != nil {
					defer db.Close()
					successFlag := new(bool)
					*successFlag = false
					c.JSON(http.StatusNotFound, responses.GamesList{Success: successFlag, Data: make([]models.GameListItem, 0), Message: err.Error()})
					return
				}

				if regionName == regionParam {
					gamePrice = price
					gameDiscount = discount
				}

				currentGameRegions = append(currentGameRegions, regionName)

				regionPriceAux := models.RegionGameListItem{
					Price:    price,
					Discount: discount,
					Region:   regionName,
				}

				currentGameRegionPrices = append(currentGameRegionPrices, regionPriceAux)

			}

			myGameAux.Prices = currentGameRegionPrices

			myGameAux.Regions = currentGameRegions

			// Mapeando desarrolladores del juego
			gameDevelopersQuery, err := db.Query("SELECT D.`idDeveloper`, D.`name`, D.`imageURL` FROM GameDeveloper AS GD INNER JOIN Developer AS D ON GD.idDeveloper = D.idDeveloper WHERE GD.idGame = ? AND GD.isDeleted = 0;", id)
			if err != nil {
				defer db.Close()
				successFlag := new(bool)
				*successFlag = false
				c.JSON(http.StatusNotFound, responses.GamesList{Success: successFlag, Data: make([]models.GameListItem, 0), Message: err.Error()})
				return
			}

			defer gameDevelopersQuery.Close()

			for gameDevelopersQuery.Next() {

				var idDeveloper int64
				var devName, devImage string

				err = gameDevelopersQuery.Scan(&idDeveloper, &devName, &devImage)
				if err != nil {
					defer db.Close()
					successFlag := new(bool)
					*successFlag = false
					c.JSON(http.StatusNotFound, responses.GamesList{Success: successFlag, Data: make([]models.GameListItem, 0), Message: err.Error()})
					return
				}

				gameDevAux := models.DeveloperGameListItem{
					DeveloperID: idDeveloper,
					Name:        devName,
					Image:       devImage,
				}

				currentGameDevelopers = append(currentGameDevelopers, gameDevAux)

			}

			myGameAux.Developer = currentGameDevelopers

			if gamePrice >= 0 && gameDiscount >= 0 {
				myGameAux.Price = gamePrice
				myGameAux.Discount = gameDiscount
				myGamesItemsList = append(myGamesItemsList, myGameAux)
			} else if isGlobal > 0 {
				myGameAux.Price = globalPrice
				myGameAux.Discount = globalDiscount
				myGamesItemsList = append(myGamesItemsList, myGameAux)
			}

		}

		successFlag := new(bool)
		*successFlag = true

		c.JSON(http.StatusOK, responses.GamesList{Success: successFlag, Data: myGamesItemsList, Message: "Info. de juegos obtenida correctamente :D"})

	}
}
