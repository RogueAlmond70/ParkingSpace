package Booking

import (
	"home/aaron/snap/go/GolandProjects/ParkingSpace/Carparks"
	"home/aaron/snap/go/GolandProjects/ParkingSpace/Users"
)

// In here we should define a booking object, and all methods etc related to booking

type Booking struct {
	User       Users.User
	Booking_ID string
	Carpark_ID string
	Carpark    Carparks.Carpark
	Space_ID   string
	Duration   int
	Brand      string
	Model      string
	Car_Size   string
}
