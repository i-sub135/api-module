package main

import (
	Ctrl "api-module/httpd/controller"
	"api-module/httpd/middleware"
	"os"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	SenDNS := os.Getenv("SENTRY_DNS")
	err = sentry.Init(sentry.ClientOptions{
		Dsn: SenDNS,
	})
	if err != nil {
		panic(err)
	}
}

func main() {
	mode := os.Getenv("MODE")
	gin.SetMode(mode)

	// initial route
	route := gin.Default()

	// include middleware
	route.Use(
		middleware.ValidREQ,
		middleware.ErrorHandle(),
	)

	// Router incude
	route.GET("/ping",
		Ctrl.IndexCtrl,
	)
	route.GET("/book",
		Ctrl.GetBookingCtrl,
	)
	route.GET("/book/:id",
		Ctrl.GetBookingUserCtrl,
	)
	route.GET("/jwt",
		Ctrl.CreateJwtCtrl,
	)

	// apps runner
	route.Run()
}
