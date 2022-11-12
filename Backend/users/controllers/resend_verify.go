package controllers

import (
	"net/http"
	"time"
	"users-service/configs"
	"users-service/responses"
	"users-service/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ResendVerify() gin.HandlerFunc {

	return func(c *gin.Context) {

		username := c.DefaultQuery("username", "Guest")
		verifyToken := uuid.New()
		emailToken := verifyToken.String()
		timeOutSecs := time.Now().Unix()
		db := configs.ConnectDB()
		defer db.Close()

		// Actualizar usuario con nuevo token de verificacion y nuevo tiempo de expiracion

		myQuery, err := db.Prepare("UPDATE `User` SET verifyToken = ?, timeOut = ? WHERE username = ?")
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.User{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		_, err = myQuery.Exec(emailToken, timeOutSecs, username)
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.User{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		// Enviar email con nuevo link de verificacion

		myQuery2, err := db.Query("SELECT email FROM User WHERE username = ?;", username)
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.User{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		var email string
		defer myQuery2.Close()
		for myQuery2.Next() {
			err = myQuery2.Scan(&email)
			if err != nil {
				defer db.Close()
				c.JSON(http.StatusBadRequest, responses.User{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
		}

		go utils.SendVerifyEmail(email, username, emailToken)

		////////////////////////////
		c.JSON(http.StatusOK, responses.User{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"resultado": "Nuevo enlace de verificacion enviado exitosamente :D"}})

	}
}
