package controller

import (
	"api-module/httpd/middleware"
	"api-module/httpd/model"
	"api-module/httpd/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// IndexCtrl -- response index
func IndexCtrl(res *gin.Context) {
	res.JSON(200, gin.H{
		"code":    200,
		"message": "ok",
	})
}

// GetBookingCtrl -- get data all booking
func GetBookingCtrl(res *gin.Context) {
	data := model.GetBooking()
	res.JSON(http.StatusOK, response.RespOK(data))
}

// GetBookingUserCtrl -- get data booking by user
func GetBookingUserCtrl(res *gin.Context) {
	id := res.Param("id")
	data := model.GetBookUser(id)

	res.JSON(http.StatusOK, response.RespOK(data))
}

// CreateJwtCtrl -- test jwt created
func CreateJwtCtrl(res *gin.Context) {
	jwt, err := middleware.CreateJwt(12, "subdiana@gmail.com")
	if err != nil {
		res.JSON(http.StatusBadRequest, response.RespBad(301, "Creating Jwt Error"))
	} else {
		res.JSON(http.StatusOK, response.RespOK(jwt))
	}
}
