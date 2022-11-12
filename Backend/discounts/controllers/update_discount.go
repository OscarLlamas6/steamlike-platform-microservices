package controllers

import (
	"net/http"

	"discounts-service/configs"
	"discounts-service/models"
	"discounts-service/responses"

	"github.com/gin-gonic/gin"
)

func UpdateDiscount() gin.HandlerFunc {

	return func(c *gin.Context) {

		var DiscountAux models.DiscountComplete
		db := configs.ConnectDB()
		defer db.Close()

		//validate the request body
		if err := c.BindJSON(&DiscountAux); err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.Discount{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&DiscountAux); validationErr != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.Discount{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		if DiscountAux.IsDLC > 0 {

			myQuery, err := db.Prepare("UPDATE `GameDiscount` SET idDLC = ?, discount = ?, startDateTime = ? , finishDateTime = ? WHERE idGameDiscount = ?")
			if err != nil {
				defer db.Close()
				c.JSON(http.StatusBadRequest, responses.Discount{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			_, err = myQuery.Exec(DiscountAux.IdDLC, DiscountAux.DiscountValue, DiscountAux.StartTime, DiscountAux.EndTime, DiscountAux.IdDiscount)
			if err != nil {
				defer db.Close()
				c.JSON(http.StatusBadRequest, responses.Discount{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}

		} else {

			myQuery, err := db.Prepare("UPDATE `GameDiscount` SET idGame = ?, discount = ?, startDateTime = ? , finishDateTime = ? WHERE idGameDiscount = ?")
			if err != nil {
				defer db.Close()
				c.JSON(http.StatusBadRequest, responses.Discount{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			_, err = myQuery.Exec(DiscountAux.IdGame, DiscountAux.DiscountValue, DiscountAux.StartTime, DiscountAux.EndTime, DiscountAux.IdDiscount)
			if err != nil {
				defer db.Close()
				c.JSON(http.StatusBadRequest, responses.Discount{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}

		}

		c.JSON(http.StatusOK, responses.Discount{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"resultado": "Descuento actualizado correctamente :D"}})
	}
}
