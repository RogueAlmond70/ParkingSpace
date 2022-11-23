package Booking

import (
	"home/aaron/snap/go/GolandProjects/ParkingSpace/Carparks"
	"home/aaron/snap/go/GolandProjects/ParkingSpace/Users"
)

// In here we should define a booking object, and all methods that act on the booking object

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

// We need storage for the bookings, in a database. See the BookingDB file for this.

// In order to make a booking, we need some kind of booking service that knows the carpark you want to use, and knows the next available space
// of the size you want to book, at the time you want to book it

func (user User) MakeBooking()
