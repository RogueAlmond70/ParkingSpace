package DBInterface

import (
	"database/sql"
	"home/aaron/snap/go/GolandProjects/ParkingSpace/Objects"
)

type UserDB interface {
	CheckIfUserExists(username string) bool
	CheckIfPasswordIsCorrect(password string) bool
}

type UserDBProd struct {
	DB sql.DB
}

type UserDBTest struct {
	DB []Objects.User
}
