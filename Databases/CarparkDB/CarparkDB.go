package CarparkDB

import (
	"home/aaron/snap/go/GolandProjects/ParkingSpace/Carparks"
)

func CreateCarpark(name string, ID string, LargeSpaces int, MediumSpaces int, SmallSpaces int) Carparks.Carpark {
	NewCarpark := Carparks.Carpark{
		Name:         name,
		ID:           ID,
		LargeSpaces:  LargeSpaces,
		MediumSpaces: MediumSpaces,
		SmallSpaces:  SmallSpaces,
	}
	return NewCarpark
}
