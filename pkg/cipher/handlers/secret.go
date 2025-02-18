package handlers

import (
	"github.com/Oxygenta-Team/FortiKey/pkg/cipher/services"

	"net/http"
)

func NewCreateSecretHandler(svc *services.SecretSvc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
