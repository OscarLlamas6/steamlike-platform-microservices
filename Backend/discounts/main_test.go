package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestAPI(t *testing.T) {
	mockResponse := `{"data":"Microservicios Descuentos - Steamlike Platform | Grupo 4 :D"}`
	r := SetUpRouter()
	r.GET("/", Saludo)
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)

	result := string(responseData)

	assert.Equal(t, mockResponse, result)
	assert.Equal(t, http.StatusOK, w.Code)
}
