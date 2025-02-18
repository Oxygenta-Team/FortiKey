package rest

import (
	"encoding/json"
	"net/http"
)

func ReturnError(w http.ResponseWriter, status int, err error) {
	ResponseJSON(w, status, &Error{
		Status:  status,
		Message: err.Error(),
	})
}

func ResponseJSON(w http.ResponseWriter, status int, data any) {
	b, err := json.Marshal(data)
	if err != nil {
		http.Error(
			w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
	}

	Respond(w, status, b)
}

func Respond(w http.ResponseWriter, code int, b []byte) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(b)
}
