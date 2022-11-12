package controllers

import (
	"net/http"
	"sales-service/configs"
	"sales-service/models"
	"sales-service/responses"
	"sales-service/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var (
	validate = validator.New()
)

func CreateSale() gin.HandlerFunc {

	return func(c *gin.Context) {

		username := c.DefaultQuery("username", "Guest")
		var newSale models.Sale
		db := configs.ConnectDB()
		defer db.Close()

		//validate the request body
		if err := c.BindJSON(&newSale); err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.Sale{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&newSale); validationErr != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.Sale{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		///////// GETTING USER EMAIL

		myQuery2, err := db.Query("SELECT email FROM User WHERE username = ?;", username)
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.Sale{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		var userEmail string = ""

		defer myQuery2.Close()

		if myQuery2.Next() {
			err = myQuery2.Scan(&userEmail)
			if err != nil {
				defer db.Close()
				c.JSON(http.StatusBadRequest, responses.Sale{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
		}

		////// REGISTRAR NUEVO SALE
		myQuery, err := db.Prepare("INSERT INTO Sale (idUser, saleDate, total, metododePago, isDeleted) VALUES(?,?,?,?,?)")
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.Sale{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		res, err := myQuery.Exec(newSale.IDUser, newSale.SaleDate, newSale.Total, newSale.MetodoDePago, 0)
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.Sale{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		lid, err := res.LastInsertId()
		if err != nil {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.Sale{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		// Convirtiendo ID de la nueva venta a string
		newSaleID := strconv.FormatInt(lid, 10)

		// Guardando todos los registros maestro-detalle de Venta-Detalle
		if !utils.SaveDetails(newSale.Detalle, newSaleID) {
			defer db.Close()
			c.JSON(http.StatusBadRequest, responses.Sale{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "Error al guardar detalles de la venta"}})
			return
		}

		go utils.SendSalesEmail(userEmail, username, newSale.SaleDate, newSale.Total, newSale.Detalle)

		c.JSON(http.StatusCreated, responses.Sale{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"resultado": "Venta registrada correctamente :D", "id": lid}})

	}
}
