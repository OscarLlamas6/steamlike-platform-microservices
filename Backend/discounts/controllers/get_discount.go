package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"discounts-service/configs"
	"discounts-service/models"
	"discounts-service/responses"

	"github.com/gin-gonic/gin"
)

func GetDiscount() gin.HandlerFunc {

	return func(c *gin.Context) {

		idDiscount := c.Param("idDiscount")
		db := configs.ConnectDB()

		myQuery, err := db.Query("SELECT * FROM GameDiscount WHERE idGameDiscount = ? AND isDeleted = 0;", idDiscount)
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.Discount{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		defer myQuery.Close()

		var DiscountAux models.DiscountComplete
		counter := 0

		for myQuery.Next() {
			var id, isDeleted, isDLC int64
			var discount float64
			var startDate, endTime string
			var idGame, idDLC sql.NullInt64

			err = myQuery.Scan(&id, &idGame, &idDLC, &discount, &startDate, &endTime, &isDeleted, &isDLC)
			if err != nil {
				defer db.Close()
				c.JSON(http.StatusBadRequest, responses.Discount{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}

			DiscountAux.IdDiscount = id

			if idGame.Valid {
				DiscountAux.IdGame = idGame.Int64
			} else {
				DiscountAux.IdGame = 0
			}

			if idDLC.Valid {
				DiscountAux.IdDLC = idDLC.Int64
			} else {
				DiscountAux.IdDLC = 0
			}

			DiscountAux.DiscountValue = discount
			DiscountAux.StartTime = startDate
			DiscountAux.EndTime = endTime
			DiscountAux.IsDLC = isDLC

			if isDeleted == 0 {
				f := new(bool)
				*f = false

				DiscountAux.IsDeleted = f
			} else {
				t := new(bool)
				*t = true
				DiscountAux.IsDeleted = t
			}

			counter++
		}

		if counter <= 0 {
			defer db.Close()
			c.JSON(http.StatusNotFound, responses.Discount{Status: http.StatusNotFound, Message: "success", Data: map[string]interface{}{"resultado": "No existe ningun descuento con ese id"}})
			return
		}

		var myDiscount map[string]interface{}
		discountJSON, err := json.Marshal(DiscountAux)
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.Discount{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		json.Unmarshal(discountJSON, &myDiscount)

		defer db.Close()
		c.JSON(http.StatusOK, responses.Discount{Status: http.StatusOK, Message: "success", Data: myDiscount})

	}
}
