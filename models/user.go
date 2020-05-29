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

func GetUsers() []*User {
	return users
}

func AddUser(u User) (User, error) {
	u.ID = nextID
	nextID++
	users = append(users, &u) // bc the user collection stores pointers
	return u, nil             // nil when there's no error
}
