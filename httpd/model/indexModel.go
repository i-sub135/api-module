package model

import (
	"api-module/httpd/db"
)

type (

	// TBooking -- table booking
	TBooking struct {
		ID             uint32
		KodeBooking    string `json:"KodeBooking,omitempty"`
		UserID         uint32
		Fullname       string
		Gender         string
		Phone          string
		Email          string
		WorkshopID     uint16
		WorkshopName   string
		CarsID         string
		PlatNomor      string
		Merek          string
		Model          string
		Odometer       uint16
		BookingDate    string
		BookingDateOld string
		Services       string
		Keluhan        string
		Cost           float64
		Paid           uint16
		PaidDate       string
		PaidChannel    string
		Created        string
		Rating         uint8
		RatingSelected string
		RatingDate     string
		RatingCount    uint8
		AlasanBatal    string
		BookingStatus  uint32
		CarsURL        string
	}
)

var sql = db.Conenect()

// GetBooking --  query get booking
func GetBooking() []TBooking {
	data := []TBooking{}
	sql.Table("t_booking").Find(&data)
	return data
}

// GetBookUser -- get booking by user
func GetBookUser(id string) []TBooking {
	data := []TBooking{}
	sql.Table("t_booking").Where("user_id=?", id).Find(&data)
	return data
}
