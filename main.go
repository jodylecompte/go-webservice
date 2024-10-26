package main

import (
	"fmt"
	"github.com/jodylecompte/go-webservice/controllers"
	"net/http"
)

func main() {
	controllers.RegisterControllers()
	err := http.ListenAndServe(":4500", nil)
	if err != nil {
		fmt.Println(err)
	}
}
