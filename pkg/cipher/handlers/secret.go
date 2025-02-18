package handlers

import (
	"encoding/json"
	"github.com/Oxygenta-Team/FortiKey/pkg/cipher/services"
	"github.com/Oxygenta-Team/FortiKey/pkg/models"
	"github.com/Oxygenta-Team/FortiKey/pkg/rest"
	"io"

	"net/http"
)

func NewCreateSecretHandler(svc services.SecretSvc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		b, err := io.ReadAll(r.Body)
		if err != nil {
			rest.ReturnError(w, http.StatusBadRequest, err)
			return
		}
		defer r.Body.Close()

		var secrets []*models.Secret
		err = json.Unmarshal(b, &secrets)
		if err != nil {
			rest.ReturnError(w, http.StatusBadRequest, err)
			return
		}

		err = svc.InsertSecret(r.Context(), secrets)
		if err != nil {
			rest.ReturnError(w, http.StatusInternalServerError, err)
			return
		}

		rest.ResponseJSON(w, http.StatusOK, secrets)
	}
}
