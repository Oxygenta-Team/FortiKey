package router

import (
	"github.com/Oxygenta-Team/FortiKey/pkg/cipher/handlers"
	"github.com/Oxygenta-Team/FortiKey/pkg/cipher/services"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(svc *services.Services) *mux.Router {
	r := mux.NewRouter()
	apiV1 := r.GetRoute("api/v1").Subrouter()

	apiV1.HandleFunc("/secret", handlers.NewCreateSecretHandler(svc.SecretSvc)).
		Methods(http.MethodPost)
	apiV1.HandleFunc("/secret", handlers.NewCompareSecretHandler(svc.SecretSvc)).
		Methods(http.MethodGet)

	return r
}
