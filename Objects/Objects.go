package Objects

type User struct {
	FirstName string
	LastName  string
	UserName  string
	Password  string
	Vehicles  []vehicle
}

// These fields will all need validation
type vehicle struct {
	Brand        string
	Model        string
	Size         string
	registration string
}
