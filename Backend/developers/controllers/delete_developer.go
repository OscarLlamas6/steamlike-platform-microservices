package controllers

import (
	"net/http"

	"developers-service/configs"
	"developers-service/responses"

	"github.com/gin-gonic/gin"
)

func DeleteDeveloper() gin.HandlerFunc {

	return func(c *gin.Context) {

		idDeveloper := c.Param("idDeveloper")
		db := configs.ConnectDB()

		myQuery, err := db.Prepare("UPDATE `Developer` SET isDeleted = 1 WHERE idDeveloper = ?")
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.Developer{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		myQuery.Exec(idDeveloper)
		defer db.Close()
		c.JSON(http.StatusOK, responses.Developer{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"resultado": "Developer eliminado correctamente :( RIP"}})
	}
}
