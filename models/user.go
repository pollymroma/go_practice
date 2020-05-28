package models

type User struct {
	ID        int
	FirstName string
	LastName  string
}

var ( // variable block
	users  []*User // a slice that has pointers to the User, a collection of addresses to users
	nextID = 1     // compiler implies type
)
