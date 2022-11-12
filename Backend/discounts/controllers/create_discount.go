package controllers

import (
	"discounts-service/configs"
	"discounts-service/models"
	"discounts-service/responses"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var (
	validate = validator.New()
)

func CreateDiscount() gin.HandlerFunc {

	return func(c *gin.Context) {

		var newDiscount models.Discount
		db := configs.ConnectDB()
		defer db.Close()

		//validate the request body
		if err := c.BindJSON(&newDiscount); err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.Discount{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&newDiscount); validationErr != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.Discount{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		////// REGISTRAR NUEVO DESCUENTO
		var newDiscountID int64

		if newDiscount.IsDLC <= 0 {

			myQuery, err := db.Prepare("INSERT INTO GameDiscount (idGame, discount, startDateTime, finishDateTime, isDeleted, isDLC) VALUES(?,?,?,?,?)")
			if err != nil {
				defer db.Close()
				c.JSON(http.StatusBadRequest, responses.Discount{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			res, err := myQuery.Exec(newDiscount.IdGame, newDiscount.DiscountValue, newDiscount.StartTime, newDiscount.EndTime, 0, 0)
			if err != nil {
				defer db.Close()
				c.JSON(http.StatusBadRequest, responses.Discount{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			newDiscountID, err = res.LastInsertId()
			if err != nil {
				defer db.Close()
				c.JSON(http.StatusBadRequest, responses.Discount{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}

		} else {

			myQuery, err := db.Prepare("INSERT INTO GameDiscount (idDLC, discount, startDateTime, finishDateTime, isDeleted, isDLC) VALUES(?,?,?,?,?)")
			if err != nil {
				defer db.Close()
				c.JSON(http.StatusBadRequest, responses.Discount{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			res, err := myQuery.Exec(newDiscount.IdDLC, newDiscount.DiscountValue, newDiscount.StartTime, newDiscount.EndTime, 0, 1)
			if err != nil {
				defer db.Close()
				c.JSON(http.StatusBadRequest, responses.Discount{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			newDiscountID, err = res.LastInsertId()
			if err != nil {
				defer db.Close()
				c.JSON(http.StatusBadRequest, responses.Discount{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}

		}

		c.JSON(http.StatusCreated, responses.Discount{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"resultado": "Descuento registrado correctamente :D", "id": newDiscountID}})

	}
}
