package controllers

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"time"
	"users-service/configs"
	"users-service/models"
	"users-service/responses"
	"users-service/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/google/uuid"
)

var (
	validate = validator.New()
)

func CreateUser() gin.HandlerFunc {

	return func(c *gin.Context) {

		var user models.User
		db := configs.ConnectDB()

		//validate the request body
		if err := c.BindJSON(&user); err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.User{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&user); validationErr != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.User{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		// Verify unique username and email

		if !IsUsernameUnique(user.UserName) {
			defer db.Close()
			c.JSON(http.StatusNonAuthoritativeInfo, responses.User{Status: http.StatusNonAuthoritativeInfo, Message: "error", Data: map[string]interface{}{"data": "Ya existe un usuario con este username"}})
			return
		}

		if !IsMailUnique(user.Email) {
			defer db.Close()
			c.JSON(http.StatusPartialContent, responses.User{Status: http.StatusPartialContent, Message: "error", Data: map[string]interface{}{"data": "Ya existe un usuario con este correo"}})
			return
		}

		///////// GETTING REGION ID

		myQuery2, err := db.Query("SELECT idRegion FROM Region WHERE name = ?;", user.Region)
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.User{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		var idRegion int64 = 1

		if myQuery2.Next() {
			err = myQuery2.Scan(&idRegion)
			if err != nil {
				defer db.Close()
				c.JSON(http.StatusBadRequest, responses.User{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
		}

		///////////////////////////

		hashString := []byte(user.Password)
		encodedPass := fmt.Sprintf("%x", md5.Sum(hashString))

		verifyToken := uuid.New()
		emailToken := verifyToken.String()
		timeOutSecs := time.Now().Unix()

		myQuery, err := db.Prepare("INSERT INTO User (name, lastName, username, birthDate, email, password, isDeleted, imageURL, isActive, timeOut, verifyToken, idRegion) VALUES(?,?,?,?,?,?,?,?,?,?,?,?)")
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.User{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		res, err := myQuery.Exec(user.Name, user.LastName, user.UserName, user.BirthDate, user.Email, encodedPass, 0, "none", 0, timeOutSecs, emailToken, idRegion)
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.User{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		lid, err := res.LastInsertId()
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.User{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		defer db.Close()

		go utils.SendVerifyEmail(user.Email, user.UserName, emailToken)

		c.JSON(http.StatusCreated, responses.User{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"resultado": "Usuario creado correctamente :D", "id": lid}})

	}
}

func IsUsernameUnique(username string) bool {

	db := configs.ConnectDB()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM User WHERE username = ?", username)
	if err != nil {
		defer db.Close()
		fmt.Println(err)
		return false
	}
	defer rows.Close()
	return !rows.Next()
}

func IsMailUnique(email string) bool {

	db := configs.ConnectDB()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM User WHERE email = ?", email)
	if err != nil {
		defer db.Close()
		fmt.Println(err)
		return false
	}
	defer rows.Close()
	return !rows.Next()
}
