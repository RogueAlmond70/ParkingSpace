package Users

import (
	"home/aaron/snap/go/GolandProjects/ParkingSpace/Booking"
	"home/aaron/snap/go/GolandProjects/ParkingSpace/Carparks"
	"home/aaron/snap/go/GolandProjects/ParkingSpace/Databases/PasswordHashing"
	"home/aaron/snap/go/GolandProjects/ParkingSpace/Vehicles"
)

type User struct {
	FirstName string
	LastName  string
	UserName  string
	Password  string
	Vehicles  []Vehicles.Vehicle
}

// Create the user and add the vehicles in two separate stages.
func CreateUser(firstname string, lastname string, username string, password string) *User {
	hashedPassword, _ := PasswordHashing.HashPassword(password)
	NewUser := &User{
		FirstName: firstname,
		LastName:  lastname,
		UserName:  username,
		Password:  hashedPassword,
		Vehicles:  nil,
	}
	return NewUser
}

func (User *User) AddVehicleToUser(brand string, model string, size string, registration string) *User {
	newVehicle := Vehicles.Vehicle{
		Brand:        brand,
		Model:        model,
		Size:         size,
		Registration: registration,
	}

	User.Vehicles = append(User.Vehicles, newVehicle)
	return User
}

// Have a think about how to assign booking IDs. It should be done sequentially
func (User User) MakeBooking() Booking.Booking {
	NewBooking := Booking.Booking{
		User:       User,
		Booking_ID: "",
		Carpark_ID: "",
		Carpark:    Carparks.Carpark{},
		Space_ID:   "",
		Duration:   0,
		Brand:      "",
		Model:      "",
		Car_Size:   "",
	}
	return NewBooking
}
