package Users

import (
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
func CreateUser(firstname string, lastname string, username string, password string) User {
	NewUser := User{
		FirstName: firstname,
		LastName:  lastname,
		UserName:  username,
		Password:  password,
		Vehicles:  nil,
	}
	return NewUser
}

func (User User) AddVehicleToUser(brand string, model string, size string, registration string) User {
	newVehicle := Vehicles.Vehicle{
		Brand:        brand,
		Model:        model,
		Size:         size,
		Registration: registration,
	}

	User.Vehicles = append(User.Vehicles, newVehicle)
	return User
}
