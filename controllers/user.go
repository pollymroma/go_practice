package controllers

import (
	"net/http"
	"regexp"
)

type userController struct {
	userIDPattern *regexp.Regexp
}

// bind function to userController type
func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// r is the request object
	w.Write([]byte("Hello from the User Controller!")) // convert string to byte slice
}

func newUserController() *userController { // returns a pointer, usually to the object
	return &userController{ //for structs you can do this, no name variable and return the address
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/'`),
	}
}
