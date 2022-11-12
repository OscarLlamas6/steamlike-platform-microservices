package controllers

import (
	"encoding/json"
	"net/http"

	"developers-service/configs"
	"developers-service/models"
	"developers-service/responses"

	"github.com/gin-gonic/gin"
)

func GetDevelopersList() gin.HandlerFunc {

	return func(c *gin.Context) {

		db := configs.ConnectDB()

		myQuery, err := db.Query("SELECT * FROM Developer WHERE isDeleted = 0")
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.Developer{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		defer myQuery.Close()

		var DeveloperAux models.DeveloperComplete
		myDevelopersList := []models.DeveloperComplete{}

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

			myDevelopersList = append(myDevelopersList, DeveloperAux)
		}

		var myFullDevelopersList []map[string]interface{}
		developerListJSON, err := json.Marshal(myDevelopersList)
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.Developer{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		json.Unmarshal(developerListJSON, &myFullDevelopersList)

		defer db.Close()
		c.JSON(http.StatusOK, responses.Developers{Status: http.StatusOK, Message: "success", Data: myFullDevelopersList})

	}
}
