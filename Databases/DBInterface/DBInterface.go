package DBInterface

import (
	"database/sql"
	"home/aaron/snap/go/GolandProjects/ParkingSpace/Objects"
)

// In here we will place our DB interface

type UserDB interface {
	checkIfUserIsValid(username string) bool
	checkIfPasswordIsCorrect(password string) bool
}

type UserDBProd struct {
	DB sql.DB
}

type UserDBTest struct {
	DB []Objects.User
}

func (UserDBTest) checkIfUserIsValid(username string) bool {
	var isValid bool
	// if user is valid, set bool to true, else set it to false
	return isValid
}

func (UserDBTest) checkIfPasswordIsCorrect(password string) bool {
	var isCorrect bool
	// if the password is correct, set bool to true, else set it to false
	return isCorrect
}

func (UserDBProd) checkIfUserIsValid(username string) bool {
	var isValid bool
	// if user is valid, set bool to true, else set it to false

	return isValid
}

func (UserDBProd) checkIfPasswordIsCorrect(password string) bool {
	var isCorrect bool
	// if the password is correct, set bool to true, else set it to false
	return isCorrect
}