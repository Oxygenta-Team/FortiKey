package router

import (
	"github.com/Oxygenta-Team/FortiKey/pkg/user-management/handler"
	"github.com/Oxygenta-Team/FortiKey/pkg/user-management/service"

	"github.com/gorilla/mux"

	"net/http"
)

func UserRoute(svs *service.UserService)  *mux.Router {
	r := mux.NewRouter()
	apiV1 := r.PathPrefix("/api/v1").Subrouter()

	apiV1.HandleFunc("/user_management", handler.NewCreateUserHandler(svs)).
		Methods(http.MethodPost)
	apiV1.HandleFunc("/user_management", handler.NewDeleteUserHandler(svs)).
		Methods(http.MethodDelete)
	apiV1.HandleFunc("/user_management", handler.NewGetUserHandler(svs)).
		Methods(http.MethodGet)

	return  r
}