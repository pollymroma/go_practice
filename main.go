package main

import (
	"net/http"

	"github.com/pluralsight/webservice/controllers"
)

func main() {
	/*u := models.User{
		ID:        2,
		FirstName: "Tricia",
		LastName:  "McMillan",
	}
	fmt.Println(u) */
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
