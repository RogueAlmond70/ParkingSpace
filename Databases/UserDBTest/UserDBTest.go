package UserDBTest

import (
	"home/aaron/snap/go/GolandProjects/ParkingSpace/Login"
	"home/aaron/snap/go/GolandProjects/ParkingSpace/Objects"
)

var TestUser Objects.User
var TestUser2 Objects.User
var AaronUser Objects.User
var Spiderman Objects.User

var TestCar Objects.Vehicle
var TestCar2 Objects.Vehicle
var SpiderCar Objects.Vehicle

var TestVehicles []Objects.Vehicle
var TestVehicles2 []Objects.Vehicle
var SpiderVehicles []Objects.Vehicle

func SetUpUserDBTest() Login.UserDBTest {
	var UserDB Login.UserDBTest

	SpiderCar.Model = "Astra"
	SpiderCar.Brand = "Vauxhall"
	SpiderCar.Size = "Medium"
	SpiderVehicles = append(SpiderVehicles, SpiderCar)

	TestCar.Model = "Leon"
	TestCar.Brand = "Seat"
	TestCar.Size = "Medium"
	TestVehicles = append(TestVehicles, TestCar)

	TestCar2.Model = "Polo"
	TestCar2.Brand = "VW"
	TestCar2.Size = "Small"
	TestVehicles2 = append(TestVehicles2, TestCar2, TestCar)

	Spiderman.UserName = "Spiderman"
	Spiderman.Password = "Manhattan"
	Spiderman.FirstName = "Peter"
	Spiderman.LastName = "Parker"
	Spiderman.Vehicles = SpiderVehicles

	TestUser.UserName = "KingAaron"
	TestUser.Password = "GobbleGobble"
	TestUser.FirstName = "Aaron"
	TestUser.LastName = "King"
	TestUser.Vehicles = TestVehicles

	TestUser2.UserName = "SantaClause"
	TestUser2.Password = "JingleJangle"
	TestUser2.FirstName = "Nicholas"
	TestUser2.LastName = "Turkey"
	TestUser2.Vehicles = TestVehicles2

	UserDB.DB = append(UserDB.DB, Spiderman, TestUser, TestUser2)

	return UserDB
}
