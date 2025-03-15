package integration_test

import (
	"bytes"
	"github.com/Oxygenta-Team/FortiKey/pkg/models"
	ta "github.com/Oxygenta-Team/FortiKey/pkg/testassets"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

func TestCreateSecretHandler(t *testing.T) {
	testTable := []struct {
		name      string
		path      string
		inputBody []*models.Secret
		//jwtToken            string TEMPORARY
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:                "OK",
			path:                "/api/v1/secrets",
			inputBody:           []*models.Secret{ta.Secret1, ta.Secret2},
			expectedStatusCode:  http.StatusOK,
			expectedRequestBody: string(ta.Marshal([]*models.Secret{ta.Secret1, ta.Secret2})),
		},
		{
			name:                "duplicate key value violates unique",
			path:                "/api/v1/secrets",
			inputBody:           []*models.Secret{ta.Secret2},
			expectedStatusCode:  http.StatusInternalServerError,
			expectedRequestBody: ta.ExpectedInternalError,
		},
	}

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodPost, ts.URL+tc.path, bytes.NewBuffer(ta.Marshal(tc.inputBody)))
			if err != nil {
				t.Fatal(err)
			}
			//req.Header.Add("Authorization", tc.jwtToken)
			res, err := http.DefaultClient.Do(req)
			defer res.Body.Close()
			body, err := io.ReadAll(res.Body)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, tc.expectedStatusCode, res.StatusCode)
			assert.Equal(t, tc.expectedRequestBody, string(body))
		})
	}
}

func TestCompareSecretHandler(t *testing.T) {
	testTable := []struct {
		name      string
		path      string
		inputBody *models.KeyValue
		//jwtToken            string   	TEMPORARY
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:                "OK",
			path:                "/api/v1/secrets",
			inputBody:           &models.KeyValue{Key: ta.Secret1.Key, Value: ta.Secret1.Value},
			expectedStatusCode:  http.StatusOK,
			expectedRequestBody: "true",
		},
		{
			name:                "Not Found",
			path:                "/api/v1/secrets",
			inputBody:           &models.KeyValue{Key: ta.Secret3.Key, Value: ta.Secret1.Value},
			expectedStatusCode:  http.StatusBadRequest,
			expectedRequestBody: ta.ExpectedNotFoundError,
		},
		{
			name:                "False Password",
			path:                "/api/v1/secrets",
			inputBody:           &models.KeyValue{Key: ta.Secret1.Key, Value: ta.Secret2.Value},
			expectedStatusCode:  http.StatusInternalServerError,
			expectedRequestBody: ta.ExpectedBcryptDecryptError,
		},
	}

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, ts.URL+tc.path, bytes.NewBuffer(ta.Marshal(tc.inputBody)))
			if err != nil {
				t.Fatal(err)
			}
			//req.Header.Add("Authorization", tc.jwtToken)
			res, err := http.DefaultClient.Do(req)
			defer res.Body.Close()
			body, err := io.ReadAll(res.Body)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, tc.expectedStatusCode, res.StatusCode)
			assert.Equal(t, tc.expectedRequestBody, string(body))
		})
	}
}
