package models

import (
	"errors"
	"fmt"
)

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
	if u.ID != 0 {
		return User{}, errors.New("New User must not include id or must be set to zero") // you MUST return a User, it sasys so in the function arguments
		//so you return an empty user
	}
	u.ID = nextID
	nextID++
	users = append(users, &u) // bc the user collection stores pointers
	return u, nil             // nil when there's no error
}

func GetUserByID(id int) (User, error) {
	for _, u := range users {
		if u.ID == id { // gets a valid user, there is no error, that's why nil
			return *u, nil
		}
	}

	return User{}, fmt.Errorf("User with ID '%v' not found", id)
}

func UpdateUser(u User) (User, error) { // loops through the collection
	for i, candidate := range users {
		if candidate.ID == u.ID {
			users[i] = &u
			return u, nil
		}
	}
	return User{}, fmt.Errorf("User with ID '%V' not found", u.ID)
}

func RemoveUserById(id int) error {
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...) // uses append to actually delete
			// like a splice, a new slice with the part before and after the found user
			return nil // returns nil to show it was successful, there's no need to return a user
		}
	}
	return fmt.Errorf("User with ID '%v' not found", id)
}
