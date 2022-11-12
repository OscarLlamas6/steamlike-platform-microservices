package controllers

import (
	"encoding/json"
	"net/http"

	"catalogs-service/configs"
	"catalogs-service/models"
	"catalogs-service/responses"

	"github.com/gin-gonic/gin"
)

func GetRegions() gin.HandlerFunc {

	return func(c *gin.Context) {

		db := configs.ConnectDB()
		defer db.Close()

		myQuery, err := db.Query("SELECT * FROM Region WHERE isDeleted = 0;")
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.Catalog{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		var RegionAux models.RegionUpdate
		myRegionsList := []models.RegionUpdate{}

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

			myRegionsList = append(myRegionsList, RegionAux)
		}

		var myFullRegionsList []map[string]interface{}
		regionsListJSON, err := json.Marshal(myRegionsList)
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.Catalog{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		json.Unmarshal(regionsListJSON, &myFullRegionsList)

		c.JSON(http.StatusOK, responses.Catalogs{Status: http.StatusOK, Message: "success", Data: myFullRegionsList})

	}
}
