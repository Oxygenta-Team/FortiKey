package handler

import (
	"encoding/json"
	"fmt"
	"github.com/Oxygenta-Team/FortiKey/pkg/models"
	"github.com/Oxygenta-Team/FortiKey/pkg/rest"
	"github.com/Oxygenta-Team/FortiKey/pkg/user-management/service"
	"net/http"
)

func NewCreateUserHandler(svc *service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			rest.ReturnError(w, http.StatusBadRequest, fmt.Errorf("Error decoding request body: %v", err))
			return
		}

		newUser, err := svc.CreateUser(user.Login,user.Email)
		if err != nil {
			rest.ReturnError(w, http.StatusInternalServerError, fmt.Errorf("error create user : %v", err))
			return
		}

		rest.ResponseJSON(w, http.StatusCreated, newUser)
	}
}

func NewDeleteUserHandler(svs *service.UserService)  http.HandlerFunc  {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if id == "" {
			rest.ReturnError(w, http.StatusBadRequest, fmt.Errorf("ID is required"))
			return
		}

		userID := 0
		if _, err := fmt.Sscanf(id, "%d", &userID); err != nil {
			rest.ReturnError(w, http.StatusBadRequest, fmt.Errorf("Invalid ID forman"))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("User deleted"))
	}
}

func NewGetUserHandler(svc *service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if id == "" {
			rest.ReturnError(w, http.StatusBadRequest, fmt.Errorf("ID is required"))
			return
		}

		userID := 0
		if _, err := fmt.Sscanf(id, "%d", &userID); err != nil {
			rest.ReturnError(w, http.StatusBadRequest, fmt.Errorf("Invalid ID format"))
			return
		}

		user, err := svc.GetUserByID(userID)
		if err != nil {
			rest.ReturnError(w, http.StatusInternalServerError, fmt.Errorf("Error getting user: %v", err))
			return
		}

		rest.ResponseJSON(w, http.StatusOK, user)
	}
}