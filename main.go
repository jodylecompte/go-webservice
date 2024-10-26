package main

import (
	"fmt"
	"github.com/jodylecompte/go-webservice/models"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Jody",
		LastName:  "LeCompte",
	}

	fmt.Println(u)
}
