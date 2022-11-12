package controllers

import (
	"encoding/json"
	"net/http"

	"catalogs-service/configs"
	"catalogs-service/models"
	"catalogs-service/responses"

	"github.com/gin-gonic/gin"
)

func GetCategory() gin.HandlerFunc {

	return func(c *gin.Context) {

		idCategory := c.Param("idCategory")
		db := configs.ConnectDB()
		defer db.Close()

		myQuery, err := db.Query("SELECT * FROM Category WHERE idCategory = ?", idCategory)
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.Catalog{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		var CategoryAux models.CategoryUpdate
		counter := 0

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

			counter++
		}

		if counter <= 0 {
			defer db.Close()
			c.JSON(http.StatusNotFound, responses.Catalog{Status: http.StatusNotFound, Message: "success", Data: map[string]interface{}{"resultado": "No existe ninguna categoria con ese id"}})
			return
		}

		var myCategory map[string]interface{}
		categoryJSON, err := json.Marshal(CategoryAux)
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.Catalog{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		json.Unmarshal(categoryJSON, &myCategory)

		c.JSON(http.StatusOK, responses.Catalog{Status: http.StatusOK, Message: "success", Data: myCategory})

	}
}
