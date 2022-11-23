package UserDB_Interface

import (
	"database/sql"
	"home/aaron/snap/go/GolandProjects/ParkingSpace/Users"
)

type UserDB interface {
	CheckIfUserExists(username string) bool
	CheckIfPasswordIsCorrect(password string) bool
}

type UserDBProd struct {
	DB sql.DB
}

type UserDBTest struct {
	DB []Users.User
}
