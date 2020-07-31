package main

import (
	_ "docs"
	"net/http"

	"github.com/labstack/echo"
	echoSwagger "github.com/swaggo/echo-swagger"
)

//SERVICE Path COUPON.
const SERVICE = "/coupon/"

// @title API Coupon MLA
// @description API to get the items that a customer can buy a coupon.
// @version 1.0
// @host localhost:8084
func main() {
	server := echo.New()

	server.GET("/", Index)
	server.POST(SERVICE, postService)
	server.GET("/swagger/*", echoSwagger.WrapHandler)
	server.Logger.Fatal(server.Start(":8084"))
}

//Index API.
func Index(server echo.Context) (err error) {
	result := `
	<h1>API CUPON - MercadoLibre</h1>
	<p>Puede dirigirse a la documentacion de la API en la siguiente URL:</p>
	<a href="http://localhost:8084/swagger/index.html">Coupon API swagger!</a>
	<p>O consumir el servicio 'http://localhost:8084/coupon/', usando el metodo POST y enviando los valores de ejemplo:</p>
	<i>{"amount": 2500,
  		"item_ids": ["MLA710902496", "MLA739047002", "MLA621847666"]}</i>`
	return server.HTML(http.StatusOK, result)
}

// postService godoc
// @Summary Get list of items
// @Description Get list of items
// @Tags Coupon
// @Accept json
// @Produce json
// @Param body body string true "Body: {item_ids, amount}"
// @Success 200 {object} response
// @Failure 400 {object} response
// @Router /coupon/ [post]
func postService(server echo.Context) (err error) {
	//Obtiene los parametros de la solicitud y se mapean al Struct 'Body'.
	items := &body{}
	if err = server.Bind(items); err != nil {
		return server.String(http.StatusBadRequest, err.Error())
	}

	if len(items.ItemIds) < 1 {
		return server.String(http.StatusBadRequest, "item_ids invalid format.")
	}

	status, finalValue := getValuesCoupon(items)
	if finalValue == nil {
		return server.String(http.StatusBadRequest, "The items do not have a valid format or price in the API MLA.")
	}

	if status == http.StatusOK {
		if finalValue.Total == 0 {
			return server.String(http.StatusNotFound, "Insufficient amount to buy an item.")
		}

		return server.JSON(http.StatusOK, finalValue)
	}

	return server.String(http.StatusBadRequest, "ERROR")
}
