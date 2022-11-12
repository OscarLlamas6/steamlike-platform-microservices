package controllers

import (
	"net/http"
	"time"

	"users-service/configs"
	"users-service/responses"

	"github.com/gin-gonic/gin"
)

func VerifyUser() gin.HandlerFunc {

	return func(c *gin.Context) {

		username := c.DefaultQuery("username", "Guest")
		verifyToken := c.DefaultQuery("verifyToken", "Guest")
		db := configs.ConnectDB()
		defer db.Close()

		rows, err := db.Query("SELECT * FROM User WHERE username = ? AND verifyToken = ?;", username, verifyToken)
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.User{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		defer rows.Close()

		if !rows.Next() {
			defer db.Close()
			c.JSON(http.StatusNotFound, responses.User{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "Enlace de verificación inválido"}})
			return
		}

		rows2, err := db.Query("SELECT * FROM User WHERE username = ? AND verifyToken = ? AND `User`.isActive = 1;", username, verifyToken)
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.User{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		defer rows2.Close()

		if rows2.Next() {
			defer db.Close()
			c.JSON(http.StatusNotModified, responses.User{Status: http.StatusNotModified, Message: "error", Data: map[string]interface{}{"data": "La cuenta ya ha sido verificada :D"}})
			return
		}

		// =============== VERIFICAR EXPIRACION DEL TOKEN ===============

		myQuery, err := db.Query("SELECT timeOut FROM User WHERE username = ?;", username)
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.User{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		defer myQuery.Close()

		var timeOut int64
		for myQuery.Next() {
			err = myQuery.Scan(&timeOut)
			if err != nil {
				defer db.Close()
				c.JSON(http.StatusBadRequest, responses.User{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
		}

		currentSeconds := time.Now().Unix()
		timeDiff := currentSeconds - timeOut

		if timeDiff >= 120 {
			defer db.Close()
			c.JSON(http.StatusGone, responses.User{Status: http.StatusGone, Message: "error", Data: map[string]interface{}{"data": "El enlace de verificacion ha expirado, presione reenviar para generar uno nuevo :D"}})
			return
		}

		// ====== FIN DE LA VERIFICACION DE LA EXPIRACION DEL TOKEN ======

		myQuery3, err := db.Prepare("UPDATE `User` SET isActive = 1 WHERE username = ?")
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.User{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		_, err = myQuery3.Exec(username)
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.User{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.User{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"resultado": "Cuenta verificada exitosamente :D"}})
	}
}
