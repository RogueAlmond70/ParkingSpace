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
