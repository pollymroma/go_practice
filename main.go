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
	println("imprime algo en consola?")
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
