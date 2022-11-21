package Objects

type User struct {
	FirstName string
	LastName  string
	UserName  string
	Password  string
	Vehicles  []Vehicle
}

// These fields will all need validation
type Vehicle struct {
	Brand        string
	Model        string
	Size         string
	registration string
}
