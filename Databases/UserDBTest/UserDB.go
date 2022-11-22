package UserDBTest

import (
	"database/sql"
	"home/aaron/snap/go/GolandProjects/ParkingSpace/Constants"
	"home/aaron/snap/go/GolandProjects/ParkingSpace/Databases/PasswordHashing"
	"home/aaron/snap/go/GolandProjects/ParkingSpace/Objects"
)

// We need a login system that stores the users basic details, including their registration number(s) and the size/type of car.
// Functional requirements
// 1. House a relational database of users with relevant info
// 2. Allow users to log in to the system and pull up their details

// Non-Functional requirements
// 1. Passwords must be hashed
// 2. The system must be highly performant

// Will need a relational database to store the users, and link to however many vehicles they own

func login(UserName string, Password string) bool {
	// If credentials are valid, return true, else return false
	var successfulLogin bool

	return successfulLogin
}

func isUserNameValid(Username string) bool {
	var isValid bool
	// Take the username, search the database with a binary search, return true if present, false if not.

	return isValid
}

type UserDBProd struct {
	DB sql.DB
}

type UserDBTest struct {
	DB []Objects.User
}

func (DB UserDBTest) CheckIfUserExists(username string) bool {
	UserDB := DB.DB
	for _, i := range UserDB {
		if username == i.UserName {
			return true
		}
	}
	return false
}

func (DB UserDBTest) CheckIfPasswordIsCorrect(username string, password string) bool {
	UserDB := DB.DB
	for _, user := range UserDB {
		if username == user.UserName {
			return PasswordHashing.CheckPasswordHash(password, Constants.Hash)
		}
	}
	return false
}

func (DB UserDBProd) CheckIfPasswordIsCorrect(username string, password string) bool {

}

func (UserDBProd) CheckIfUserExists(username string) bool {
	var isValid bool
	// if user is valid, set bool to true, else set it to false

	return isValid
}
