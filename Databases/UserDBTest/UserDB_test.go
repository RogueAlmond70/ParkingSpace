package UserDBTest

import (
	"home/aaron/snap/go/GolandProjects/ParkingSpace/Objects"
	"testing"
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

func SetUpUserDBTest() UserDBTest {
	var UserDB UserDBTest

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

func TestCheckIfUserIsValid(t *testing.T) {
	TestDB := SetUpUserDBTest()
	var tests = []struct {
		actual   bool
		expected bool
	}{
		{TestDB.CheckIfUserExists("KingAaron"), true},
		{TestDB.CheckIfUserExists("testUser"), false},
		{TestDB.CheckIfUserExists("Spiderman"), true},
		{TestDB.CheckIfUserExists("SPiDerMan"), false},
		{TestDB.CheckIfUserExists("SantaClause"), true},
		{TestDB.CheckIfUserExists("SantaClause1"), false},
	}
	for _, test := range tests {
		if output := test.actual; output != test.expected {
			t.Errorf("Test Failed: expected #{test.expected}, received #{test.actual}")
		}
	}
}
