package main

import (
	"fmt"
	"github.com/jodylecompte/go-webservice/controllers"
	"net/http"
)

func main() {
	controllers.RegisterControllers()

	fmt.Println("Listening on port 4500...")

	err := http.ListenAndServe(":4500", nil)
	if err != nil {
		fmt.Println("Server failed to start:", err)
	}
}
