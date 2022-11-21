package DBInterface

// In here we will place our DB interface

type UserDB interface {
	checkIfUserIsValid(username string) bool
	checkIfPasswordIsCorrect(password string) bool
}
