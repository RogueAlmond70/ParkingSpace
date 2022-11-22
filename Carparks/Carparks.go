package Carparks

// This should house the carparks struct, and any relevant methods or functions

type Carpark struct {
	Name         string
	ID           string
	TotalSpaces  int
	LargeSpaces  int
	MediumSpaces int
	SmallSpaces  int
}

// Decide if you want to have a database for this or just use in memory storage like a slice.
var TestCarparks []Carpark

// ProdCarparks will be a database table

func AddCarPark()
