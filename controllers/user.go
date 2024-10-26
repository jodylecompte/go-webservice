package controllers

import (
	"net/http"
	"regexp"
)

type userController struct {
	userIDPattern *regexp.Regexp
}

func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello from the user controller"))
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
}

func newUserController() *userController {
	return &userController{
		userIDPattern: regexp.MustCompile("^/users/([0-9]+)"),
	}
}
