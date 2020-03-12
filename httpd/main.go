package main

import (
	Ctrl "api-module/httpd/controller"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func main() {

	mode := os.Getenv("MODE")
	gin.SetMode(mode)

	route := gin.Default()
	route.GET("/ping", Ctrl.IndexCtrl)
	route.GET("/book", Ctrl.GetBookingCtrl)
	route.GET("/book/:id", Ctrl.GetBookingUserCtrl)

	route.Run()
}
