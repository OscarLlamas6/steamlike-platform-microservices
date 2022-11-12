package controllers

import (
	"encoding/json"
	"net/http"

	"users-service/configs"
	"users-service/models"
	"users-service/responses"

	"github.com/gin-gonic/gin"
)

func GetProfile() gin.HandlerFunc {

	return func(c *gin.Context) {

		userName := c.Param("username")
		db := configs.ConnectDB()

		myQuery, err := db.Query("SELECT * FROM User WHERE username = ?", userName)
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.User{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		var User models.UserComplete
		counter := 0

		defer myQuery.Close()

		for myQuery.Next() {
			var id, isDeleted, isActive, timeOut, idRegion int64

			var name, lastName, username, birthDate, email, password, imageURL, verifyToken string

			err = myQuery.Scan(&id, &name, &lastName, &username, &birthDate, &email, &password, &isDeleted, &imageURL, &isActive, &verifyToken, &idRegion, &timeOut)
			if err != nil {
				defer db.Close()
				c.JSON(http.StatusBadRequest, responses.User{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}

			User.Id = id
			User.Name = name
			User.LastName = lastName
			User.UserName = username
			User.BirthDate = birthDate
			User.Email = email
			User.Password = password
			User.Region = idRegion

			if isDeleted == 0 {
				f := new(bool)
				*f = false

				User.IsDeleted = f
			} else {
				t := new(bool)
				*t = true
				User.IsDeleted = t
			}

			User.ImageURL = imageURL

			if isActive == 0 {
				f := new(bool)
				*f = false
				User.IsActive = f
			} else {
				t := new(bool)
				*t = true
				User.IsActive = t
			}

			User.VerifyToken = verifyToken
			counter++
		}

		if counter <= 0 {
			defer db.Close()
			c.JSON(http.StatusOK, responses.User{Status: http.StatusNotFound, Message: "success", Data: map[string]interface{}{"resultado": "No existe ningun usuario con ese id"}})
			return
		}

		var myUser map[string]interface{}
		UserJson, _ := json.Marshal(User)
		json.Unmarshal(UserJson, &myUser)

		defer db.Close()
		c.JSON(http.StatusOK, responses.User{Status: http.StatusOK, Message: "success", Data: myUser})

	}
}
