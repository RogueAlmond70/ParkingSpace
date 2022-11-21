package UserDBTest

import (
	"home/aaron/snap/go/GolandProjects/ParkingSpace/Databases/DBInterface"
	"home/aaron/snap/go/GolandProjects/ParkingSpace/Objects"
)

var Spiderman Objects.User
var TestUser Objects.User


func SetUpUserDBTest() DBInterface.UserDBTest {
	var UserSlice []Objects.User
	UserSlice = append(UserSlice,{ "Peter", "Parker", "Spiderman", "Manhattan"})

	UserDBTest := DBInterface.UserDBTest{
		{{"Peter", "Parker", "Spiderman", "Manhattan"}},
	}
}
