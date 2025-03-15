package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
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

		err = svc.CreateSecret(r.Context(), secrets)
		if err != nil {
			rest.ReturnError(w, http.StatusInternalServerError, services.ErrInternal)
			return
		}

		rest.ResponseJSON(w, http.StatusOK, secrets)
	}
}

func NewCompareSecretHandler(svc services.SecretSvc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		b, err := io.ReadAll(r.Body)
		if err != nil {
			rest.ReturnError(w, http.StatusBadRequest, err)
			return
		}
		defer r.Body.Close()

		var keyValue *models.KeyValue
		err = json.Unmarshal(b, &keyValue)
		if err != nil {
			rest.ReturnError(w, http.StatusBadRequest, err)
			return
		}

		compare, err := svc.CompareSecret(r.Context(), keyValue)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				rest.ReturnError(w, http.StatusBadRequest, services.ErrNotFound)
				return
			}
			rest.ReturnError(w, http.StatusInternalServerError, err)
			return
		}

		rest.ResponseJSON(w, http.StatusOK, compare)
	}
}
