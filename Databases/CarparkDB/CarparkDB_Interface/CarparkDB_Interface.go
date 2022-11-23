package CarparkDB_Interface

import (
	"database/sql"
	"home/aaron/snap/go/GolandProjects/ParkingSpace/Carparks"
)

type CarparkDB interface {
	AddCarpark(carpark Carparks.Carpark)
}

type CarparkDBTest struct {
	DB sql.DB
}

func (CarparkDB) AddCarpark(carpark Carparks.Carpark) {

}
