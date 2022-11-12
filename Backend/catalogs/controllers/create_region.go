package controllers

import (
	"catalogs-service/configs"
	"catalogs-service/models"
	"catalogs-service/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateRegion() gin.HandlerFunc {

	return func(c *gin.Context) {

		var region models.Region
		db := configs.ConnectDB()
		defer db.Close()

		//validate the request body
		if err := c.BindJSON(&region); err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.Catalog{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&region); validationErr != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.Catalog{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		////// AGREGAR CATEGORIA NUEVA

		myQuery, err := db.Prepare("INSERT INTO Region (name, isDeleted) VALUES(?,?)")
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.Catalog{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		res, err := myQuery.Exec(region.Name, 0)
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.Catalog{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		lid, err := res.LastInsertId()
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.Catalog{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.Catalog{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"resultado": "Nueva region agregada correctamente :D", "id": lid}})

	}
}
