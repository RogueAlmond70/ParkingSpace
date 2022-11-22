package Users

import "home/aaron/snap/go/GolandProjects/ParkingSpace/Objects"

type User struct {
	FirstName string
	LastName  string
	UserName  string
	Password  string
	Vehicles  []Objects.Vehicle
}

// We want all the methods relating to the user in here
