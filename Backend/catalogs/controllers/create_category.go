package controllers

import (
	"catalogs-service/configs"
	"catalogs-service/models"
	"catalogs-service/responses"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var (
	validate = validator.New()
)

func CreateCategory() gin.HandlerFunc {

	return func(c *gin.Context) {

		var category models.Category
		db := configs.ConnectDB()

		//validate the request body
		if err := c.BindJSON(&category); err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.Catalog{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&category); validationErr != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.Catalog{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		////// AGREGAR CATEGORIA NUEVA

		myQuery, err := db.Prepare("INSERT INTO Category (name, isDeleted) VALUES(?,?)")
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.Catalog{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		res, err := myQuery.Exec(category.Name, 0)
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
		defer db.Close()

		c.JSON(http.StatusCreated, responses.Catalog{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"resultado": "Nueva categoria agregada correctamente :D", "id": lid}})

	}
}
