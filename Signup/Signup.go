package Signup

import (
	"home/aaron/snap/go/GolandProjects/ParkingSpace/Databases/DBInterface"
	"home/aaron/snap/go/GolandProjects/ParkingSpace/Objects"
)

func (DB DBInterface.UserDB) Signup(firstname string, lastname string, username string, password string, vehicles []Objects.Vehicle) {

}

func (Objects.User)Signup(DB DBInterface.UserDB)error{

}

func