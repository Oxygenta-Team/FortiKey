package rest

import (
	"encoding/json"
	"net/http"
)

func ReturnError(w http.ResponseWriter, err error, status int) {
	error := writeJSON(w, err, status)
	if error != nil {
		w.Write([]byte(err.Error()))
		return
	}
}

func writeJSON(w http.ResponseWriter, data any, status int) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	w.WriteHeader(status)
	_, err = w.Write(b)
	return err
}
