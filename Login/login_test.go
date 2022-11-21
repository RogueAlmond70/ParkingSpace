package Login

import (
	UserDBTest2 "home/aaron/snap/go/GolandProjects/ParkingSpace/Databases/UserDBTest"
	"testing"
)

func TestCheckIfUserIsValid(t *testing.T) {
	TestDB := UserDBTest2.SetUpUserDBTest()
	var tests = []struct {
		actual   bool
		expected bool
	}{
		{TestDB.CheckIfUserIsValid("KingAaron"), true},
		{TestDB.CheckIfUserIsValid("testUser"), false},
		{TestDB.CheckIfUserIsValid("Spiderman"), true},
		{TestDB.CheckIfUserIsValid("SPiDerMan"), false},
		{TestDB.CheckIfUserIsValid("SantaClause"), true},
		{TestDB.CheckIfUserIsValid("SantaClause1"), false},
	}
	for _, test := range tests {
		if output := test.actual; output != test.expected {
			t.Errorf("Test Failed: expected #{test.expected}, received #{test.actual}")
		}
	}
}
