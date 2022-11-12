package controllers

import (
	"crypto/md5"
	"fmt"
	"net/http"

	"users-service/configs"
	"users-service/models"
	"users-service/responses"

	"github.com/gin-gonic/gin"
)

func UpdateUser() gin.HandlerFunc {

	return func(c *gin.Context) {

		var User models.UserUpdate
		db := configs.ConnectDB()

		//validate the request body
		if err := c.BindJSON(&User); err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.User{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&User); validationErr != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.User{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		///////// CHECK UNIQUE EMAIL

		myQuery2, err := db.Query("SELECT * FROM User WHERE email = ? AND username <> ?;", User.Email, User.UserName)
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.User{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		defer myQuery2.Close()

		if myQuery2.Next() {
			if err != nil {
				defer db.Close()
				c.JSON(http.StatusPartialContent, responses.User{Status: http.StatusPartialContent, Message: "error", Data: map[string]interface{}{"data": "Ya existe otro usuario con este correo"}})
				return
			}
		}

		var updatePass string = ""

		////// CHECK FOR NEW PASSWOROD

		userID := User.UserName
		hashString := []byte(User.OldPassword)
		encodedPass := fmt.Sprintf("%x", md5.Sum(hashString))

		// Verificando si las credenciales coinciden con el usuario correcto

		userQuery, err := db.Query("SELECT * FROM User WHERE username = ? AND password = ?", userID, encodedPass)
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.User{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		defer userQuery.Close()

		if !userQuery.Next() {
			defer db.Close()
			c.JSON(http.StatusNotAcceptable, responses.User{Status: http.StatusNotAcceptable, Message: "error", Data: map[string]interface{}{"resultado": "El password actual proporcionado no coincide"}})
			return
		} else {

			if User.SetNewPass > 0 {
				hashString := []byte(User.NewPassword)
				encodedPass := fmt.Sprintf("%x", md5.Sum(hashString))
				updatePass = encodedPass
			} else {
				hashString := []byte(User.OldPassword)
				encodedPass := fmt.Sprintf("%x", md5.Sum(hashString))
				updatePass = encodedPass
			}

		}

		myQuery, err := db.Prepare("UPDATE `User` SET name = ?, lastName = ?, birthDate = ?, email = ?, password = ?, idRegion = ? WHERE idUser = ?")
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.User{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		_, err = myQuery.Exec(User.Name, User.LastName, User.BirthDate, User.Email, updatePass, User.Region, User.Id)
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.User{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		defer db.Close()
		c.JSON(http.StatusOK, responses.User{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"resultado": "Usuario actualizado correctamente :D"}})
	}
}
