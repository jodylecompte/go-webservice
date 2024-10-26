package controllers

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"

	"github.com/jodylecompte/go-webservice/models"
)

type userController struct {
	userIDPattern *regexp.Regexp
}

func (uc *userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/users" || r.URL.Path == "/users/" {
		switch r.Method {
		case http.MethodGet:
			uc.getAll(w, r)
		case http.MethodPost:
			uc.post(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	} else {
		matches := uc.userIDPattern.FindStringSubmatch(r.URL.Path)

		if len(matches) == 0 {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}

		id, err := strconv.Atoi(matches[1])
		if err != nil {
			http.Error(w, "Invalid User ID", http.StatusBadRequest)
			return
		}

		switch r.Method {
		case http.MethodGet:
			uc.get(id, w)
		case http.MethodPut:
			uc.put(id, w, r)
		case http.MethodDelete:
			uc.delete(id, w)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	}
}

func (uc *userController) getAll(w http.ResponseWriter, r *http.Request) {
	users := models.GetUsers()
	if err := encodeResponseAsJSON(users, w); err != nil {
		http.Error(w, "Failed to encode users", http.StatusInternalServerError)
	}
}

func (uc *userController) get(id int, w http.ResponseWriter) {
	user, err := models.GetUserByID(id)
	if err != nil {
		http.Error(w, "User Not Found", http.StatusNotFound)
		return
	}
	if err := encodeResponseAsJSON(user, w); err != nil {
		http.Error(w, "Failed to encode user", http.StatusInternalServerError)
	}
}

func (uc *userController) post(w http.ResponseWriter, r *http.Request) {
	user, err := uc.parseRequest(r)
	if err != nil {
		http.Error(w, "Could not parse User object", http.StatusBadRequest)
		return
	}
	user, err = models.AddUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := encodeResponseAsJSON(user, w); err != nil {
		http.Error(w, "Failed to encode user", http.StatusInternalServerError)
	}
}

func (uc *userController) put(id int, w http.ResponseWriter, r *http.Request) {
	user, err := uc.parseRequest(r)

	if err != nil {
		http.Error(w, "Could not parse User object", http.StatusBadRequest)
		return
	}

	user.ID = id
	user, err = models.UpdateUser(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := encodeResponseAsJSON(user, w); err != nil {
		http.Error(w, "Failed to encode user", http.StatusInternalServerError)
	}
}

func (uc *userController) delete(id int, w http.ResponseWriter) {
	err := models.RemoveUserByID(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (uc *userController) parseRequest(r *http.Request) (models.User, error) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		return models.User{}, err
	}
	return user, nil
}

func encodeResponseAsJSON(data interface{}, w http.ResponseWriter) error {
	if err := json.NewEncoder(w).Encode(data); err != nil {
		return err
	}
	return nil
}

func newUserController() *userController {
	return &userController{
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?$`),
	}
}
