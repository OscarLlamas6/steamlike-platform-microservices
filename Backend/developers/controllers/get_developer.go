package controllers

import (
	"encoding/json"
	"net/http"

	"developers-service/configs"
	"developers-service/models"
	"developers-service/responses"

	"github.com/gin-gonic/gin"
)

func GetDeveloper() gin.HandlerFunc {

	return func(c *gin.Context) {

		idDeveloper := c.Param("idDeveloper")
		db := configs.ConnectDB()

		myQuery, err := db.Query("SELECT * FROM Developer WHERE idDeveloper = ? AND isDeleted = 0", idDeveloper)
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.Developer{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		var DeveloperAux models.DeveloperComplete
		counter := 0

		defer myQuery.Close()

		for myQuery.Next() {
			var id, isDeleted int64
			var name, country, imageURL, email string

			err = myQuery.Scan(&id, &name, &country, &imageURL, &email, &isDeleted)
			if err != nil {
				defer db.Close()
				c.JSON(http.StatusBadRequest, responses.Developer{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}

			DeveloperAux.IdDeveloper = id
			DeveloperAux.Name = name
			DeveloperAux.Pais = country
			DeveloperAux.ImageURL = imageURL
			DeveloperAux.Email = email

			if isDeleted == 0 {
				f := new(bool)
				*f = false

				DeveloperAux.IsDeleted = f
			} else {
				t := new(bool)
				*t = true
				DeveloperAux.IsDeleted = t
			}

			counter++
		}

		if counter <= 0 {
			defer db.Close()
			c.JSON(http.StatusNotFound, responses.Developer{Status: http.StatusNotFound, Message: "success", Data: map[string]interface{}{"resultado": "No existe ningun developer con ese id"}})
			return
		}

		var myDeveloper map[string]interface{}
		developerJSON, err := json.Marshal(DeveloperAux)
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.Developer{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		json.Unmarshal(developerJSON, &myDeveloper)

		defer db.Close()
		c.JSON(http.StatusOK, responses.Developer{Status: http.StatusOK, Message: "success", Data: myDeveloper})

	}
}
