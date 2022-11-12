package controllers

import (
	"encoding/json"
	"net/http"

	"catalogs-service/configs"
	"catalogs-service/models"
	"catalogs-service/responses"

	"github.com/gin-gonic/gin"
)

func GetCategories() gin.HandlerFunc {

	return func(c *gin.Context) {

		db := configs.ConnectDB()
		defer db.Close()

		myQuery, err := db.Query("SELECT * FROM Category WHERE isDeleted = 0;")
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.Catalog{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		var CategoryAux models.CategoryUpdate
		myCategoriesList := []models.CategoryUpdate{}

		defer myQuery.Close()

		for myQuery.Next() {
			var id, isDeleted int64

			var name string

			err = myQuery.Scan(&id, &name, &isDeleted)
			if err != nil {
				defer db.Close()
				c.JSON(http.StatusBadRequest, responses.Catalog{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}

			CategoryAux.IDCategory = id
			CategoryAux.Name = name

			if isDeleted == 0 {
				f := new(bool)
				*f = false

				CategoryAux.IsDeleted = f
			} else {
				t := new(bool)
				*t = true
				CategoryAux.IsDeleted = t
			}

			myCategoriesList = append(myCategoriesList, CategoryAux)
		}

		var myFullCategoriesList []map[string]interface{}
		categoryListJSON, err := json.Marshal(myCategoriesList)
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.Catalog{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		json.Unmarshal(categoryListJSON, &myFullCategoriesList)

		c.JSON(http.StatusOK, responses.Catalogs{Status: http.StatusOK, Message: "success", Data: myFullCategoriesList})

	}
}
