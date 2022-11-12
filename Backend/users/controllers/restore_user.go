package controllers

import (
	"net/http"

	"users-service/configs"
	"users-service/responses"

	"github.com/gin-gonic/gin"
)

func RestoreUser() gin.HandlerFunc {

	return func(c *gin.Context) {

		username := c.Param("username")
		db := configs.ConnectDB()

		myQuery, err := db.Prepare("UPDATE `User` SET isDeleted = 0 WHERE username = ?")
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.User{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		myQuery.Exec(username)
		defer db.Close()
		c.JSON(http.StatusOK, responses.User{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"resultado": "Usuario restaurado correctamente :D"}})
	}
}
