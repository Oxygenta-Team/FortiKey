package testassets

import (
	"github.com/Oxygenta-Team/FortiKey/pkg/models"
)

var (
	ID500 = uint64(500)
	ID501 = uint64(501)
	ID502 = uint64(502)

	Value1 = "secretPassword"
	Value2 = "secretPassword_2"

	Hash1 = Marshal(Value1)
	Hash2 = Marshal(Value2)
)

var (
	Secret1 = &models.Secret{
		ID:     ID500,
		UserID: ID500,
		Key:    "user_name_1",
		Value:  Value1,
		Method: models.BCRYPT,
		Hash:   Hash1,
	}
	Secret2 = &models.Secret{
		ID:     ID501,
		UserID: ID501,
		Key:    "user_name_2",
		Value:  Value2,
		Method: models.BCRYPT,
		Hash:   Hash2,
	}
	Secret3 = &models.Secret{
		ID:     ID502,
		UserID: ID502,
		Key:    "user_name_3",
		Value:  Value2,
		Method: models.BCRYPT,
		Hash:   Hash2,
	}
)
