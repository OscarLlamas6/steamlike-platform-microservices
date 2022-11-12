package controllers

import (
	"games-service/configs"
	"games-service/models"
	"games-service/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetGame() gin.HandlerFunc {

	return func(c *gin.Context) {

		idGame := c.Param("idGame")
		db := configs.ConnectDB()
		defer db.Close()

		gameListQuery, err := db.Query("SELECT * FROM Game WHERE idGame = ? AND isDeleted = 0;", idGame)
		if err != nil {
			defer db.Close()
			successFlag := new(bool)
			*successFlag = false
			c.JSON(http.StatusNotFound, responses.GameInfo{Success: successFlag, Data: models.FullGameListItem{}, Message: err.Error()})
			return
		}

		// Struct para almacenar toda la info del juego
		myGameAux := models.FullGameListItem{}
		counter := 0

		defer gameListQuery.Close()

		for gameListQuery.Next() {

			// Arreglo para almacenar las categorias del juego
			currentGameCategories := []models.CategoryItem{}

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
				c.JSON(http.StatusNotFound, responses.GameInfo{Success: successFlag, Data: models.FullGameListItem{}, Message: err.Error()})
				return
			}

			myGameAux.GameID = id
			myGameAux.Name = name
			myGameAux.Image = imageURL
			myGameAux.ReleaseDate = releaseDate
			myGameAux.RestrictionAge = restrictionAge
			myGameAux.Description = description
			myGameAux.Group = group
			myGameAux.Price = globalPrice
			myGameAux.Discount = globalDiscount

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
			gameCategoryQuery, err := db.Query("SELECT `Category`.`idCategory`, `Category`.`name` FROM GameCategory INNER JOIN Category ON GameCategory.idCategory = Category.idCategory WHERE `GameCategory`.`idGame` = ? AND `GameCategory`.isDeleted = 0;", id)
			if err != nil {
				defer db.Close()
				successFlag := new(bool)
				*successFlag = false
				c.JSON(http.StatusNotFound, responses.GameInfo{Success: successFlag, Data: models.FullGameListItem{}, Message: err.Error()})
				return
			}

			defer gameCategoryQuery.Close()

			for gameCategoryQuery.Next() {

				var categoryID int64
				var categoryName string

				err = gameCategoryQuery.Scan(&categoryID, &categoryName)
				if err != nil {
					defer db.Close()
					successFlag := new(bool)
					*successFlag = false
					c.JSON(http.StatusNotFound, responses.GameInfo{Success: successFlag, Data: models.FullGameListItem{}, Message: err.Error()})
					return
				}

				categoryAux := models.CategoryItem{
					CategoryID: categoryID,
					Name:       categoryName,
				}

				currentGameCategories = append(currentGameCategories, categoryAux)

			}

			myGameAux.Categories = currentGameCategories

			// Mapeando Precios por Region del juego
			pricePerRegionQuery, err := db.Query("SELECT RP.`price`, RP.`discount`, R.`name` FROM RegionPrice AS RP INNER JOIN Region AS R ON RP.idRegion = R.idRegion WHERE RP.idGame = ? AND RP.isDeleted = 0 AND RP.isDLC = 0;", id)
			if err != nil {
				defer db.Close()
				successFlag := new(bool)
				*successFlag = false
				c.JSON(http.StatusNotFound, responses.GameInfo{Success: successFlag, Data: models.FullGameListItem{}, Message: err.Error()})
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
					c.JSON(http.StatusNotFound, responses.GameInfo{Success: successFlag, Data: models.FullGameListItem{}, Message: err.Error()})
					return
				}

				regionPriceAux := models.RegionGameListItem{
					Price:    price,
					Discount: discount,
					Region:   regionName,
				}

				currentGameRegionPrices = append(currentGameRegionPrices, regionPriceAux)

			}

			myGameAux.RegionsAndPrices = currentGameRegionPrices

			// Mapeando desarrolladores del juego
			gameDevelopersQuery, err := db.Query("SELECT D.`idDeveloper`, D.`name`, D.`imageURL` FROM GameDeveloper AS GD INNER JOIN Developer AS D ON GD.idDeveloper = D.idDeveloper WHERE GD.idGame = ? AND GD.isDeleted = 0;", id)
			if err != nil {
				defer db.Close()
				successFlag := new(bool)
				*successFlag = false
				c.JSON(http.StatusNotFound, responses.GameInfo{Success: successFlag, Data: models.FullGameListItem{}, Message: err.Error()})
				return
			}

			defer gameCategoryQuery.Close()

			for gameDevelopersQuery.Next() {

				var idDeveloper int64
				var devName, devImage string

				err = gameDevelopersQuery.Scan(&idDeveloper, &devName, &devImage)
				if err != nil {
					defer db.Close()
					successFlag := new(bool)
					*successFlag = false
					c.JSON(http.StatusNotFound, responses.GameInfo{Success: successFlag, Data: models.FullGameListItem{}, Message: err.Error()})
					return
				}

				gameDevAux := models.DeveloperGameListItem{
					DeveloperID: idDeveloper,
					Name:        devName,
					Image:       devImage,
				}

				currentGameDevelopers = append(currentGameDevelopers, gameDevAux)

			}

			myGameAux.Developers = currentGameDevelopers

			counter++
		}

		successFlag := new(bool)
		*successFlag = true
		c.JSON(http.StatusOK, responses.GameInfo{Success: successFlag, Data: myGameAux, Message: "Info. de juegos obtenida correctamente :D"})

	}
}
