package controllers

import (
	"encoding/json"
	"net/http"

	"catalogs-service/configs"
	"catalogs-service/models"
	"catalogs-service/responses"

	"github.com/gin-gonic/gin"
)

func GetRegion() gin.HandlerFunc {

	return func(c *gin.Context) {

		idRegion := c.Param("idRegion")
		db := configs.ConnectDB()
		defer db.Close()

		myQuery, err := db.Query("SELECT * FROM Region WHERE idRegion = ?", idRegion)
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.Catalog{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		var RegionAux models.RegionUpdate
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

			RegionAux.IDRegion = id
			RegionAux.Name = name

			if isDeleted == 0 {
				f := new(bool)
				*f = false

				RegionAux.IsDeleted = f
			} else {
				t := new(bool)
				*t = true
				RegionAux.IsDeleted = t
			}

			counter++
		}

		if counter <= 0 {
			defer db.Close()
			c.JSON(http.StatusNotFound, responses.Catalog{Status: http.StatusNotFound, Message: "success", Data: map[string]interface{}{"resultado": "No existe ninguna region con ese id"}})
			return
		}

		var myCategory map[string]interface{}
		categoryJSON, err := json.Marshal(RegionAux)
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.Catalog{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		json.Unmarshal(categoryJSON, &myCategory)

		c.JSON(http.StatusOK, responses.Catalog{Status: http.StatusOK, Message: "success", Data: myCategory})

	}
}
