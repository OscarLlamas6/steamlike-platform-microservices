package controllers

import (
	"net/http"

	"users-service/configs"
	"users-service/models"
	"users-service/responses"

	"github.com/gin-gonic/gin"
)

func GetAllUsers() gin.HandlerFunc {

	return func(c *gin.Context) {

		db := configs.ConnectDB()
		defer db.Close()

		//Arreglo para almacenar todos los usuarios
		myUsersList := []models.UserComplete{}

		myQuery, err := db.Query("SELECT * FROM User;")
		if err != nil {
			defer db.Close()
			successFlag := new(bool)
			*successFlag = false
			c.JSON(http.StatusNotFound, responses.UsersList{Success: successFlag, Data: make([]models.UserComplete, 0), Message: err.Error()})
			return
		}

		defer myQuery.Close()

		for myQuery.Next() {

			// Struct para almacenar toda la info del juego
			myUserAux := models.UserComplete{}

			var id, isDeleted, isActive, timeOut, idRegion int64

			var name, lastName, username, birthDate, email, password, imageURL, verifyToken string

			err = myQuery.Scan(&id, &name, &lastName, &username, &birthDate, &email, &password, &isDeleted, &imageURL, &isActive, &verifyToken, &idRegion, &timeOut)
			if err != nil {
				defer db.Close()
				successFlag := new(bool)
				*successFlag = false
				c.JSON(http.StatusNotFound, responses.UsersList{Success: successFlag, Data: make([]models.UserComplete, 0), Message: err.Error()})
				return
			}

			myUserAux.Id = id
			myUserAux.Name = name
			myUserAux.LastName = lastName
			myUserAux.UserName = username
			myUserAux.BirthDate = birthDate
			myUserAux.Email = email
			myUserAux.Password = password
			myUserAux.Region = idRegion
			myUserAux.TimeOut = timeOut

			if isDeleted == 0 {
				f := new(bool)
				*f = false

				myUserAux.IsDeleted = f
			} else {
				t := new(bool)
				*t = true
				myUserAux.IsDeleted = t
			}

			myUserAux.ImageURL = imageURL

			if isActive == 0 {
				f := new(bool)
				*f = false
				myUserAux.IsActive = f
			} else {
				t := new(bool)
				*t = true
				myUserAux.IsActive = t
			}

			myUserAux.VerifyToken = verifyToken

			myUsersList = append(myUsersList, myUserAux)

		}

		successFlag := new(bool)
		*successFlag = true

		c.JSON(http.StatusOK, responses.UsersList{Success: successFlag, Data: myUsersList, Message: "Info. de usuarios obtenida correctamente :D"})

	}
}
