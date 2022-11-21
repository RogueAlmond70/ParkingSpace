package Login

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
