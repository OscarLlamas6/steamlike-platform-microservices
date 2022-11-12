package controllers

import (
	"encoding/json"
	"net/http"

	"discounts-service/configs"
	"discounts-service/models"
	"discounts-service/responses"

	"github.com/gin-gonic/gin"
)

func GetDLCDiscounts() gin.HandlerFunc {

	return func(c *gin.Context) {

		db := configs.ConnectDB()
		defer db.Close()

		myQuery, err := db.Query("SELECT idGameDiscount, idDLC, discount, startDateTime, finishDateTime, isDeleted, isDLC FROM GameDiscount WHERE isDLC = 1 AND isDeleted = 0;")
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.Discount{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		defer myQuery.Close()

		myDiscountsList := []models.DiscountComplete{}

		for myQuery.Next() {

			var DiscountAux models.DiscountComplete

			var id, idDLC, isDeleted, isDLC int64
			var discount float64
			var startTime, endTime string

			err = myQuery.Scan(&id, &idDLC, &discount, &startTime, &endTime, &isDeleted, &isDLC)
			if err != nil {
				defer db.Close()
				c.JSON(http.StatusBadRequest, responses.Discount{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}

			DiscountAux.IdDiscount = id
			DiscountAux.IdDLC = idDLC
			DiscountAux.DiscountValue = discount
			DiscountAux.StartTime = startTime
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

			myDiscountsList = append(myDiscountsList, DiscountAux)
		}

		var myFullDiscountsList []map[string]interface{}
		discountsListJSON, err := json.Marshal(myDiscountsList)
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.Discount{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		json.Unmarshal(discountsListJSON, &myFullDiscountsList)

		c.JSON(http.StatusOK, responses.Discounts{Status: http.StatusOK, Message: "success", Data: myFullDiscountsList})

	}
}
