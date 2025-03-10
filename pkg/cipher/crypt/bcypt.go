package crypt

import (
	"github.com/Oxygenta-Team/FortiKey/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

func BCryptSecret(secret *models.Secret) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(secret.Value), bcrypt.MaxCost)
	if err != nil {
		return ErrBcryptGenerate
	}
	secret.Hash = hashedPassword
	secret.Method = models.BCRYPT
	secret.UserID = 1000 // TODO Temporary, we need a user-management

	return nil
}

func BCryptCompare(hash []byte, value string) bool {
	err := bcrypt.CompareHashAndPassword(hash, []byte(value))

	return err == nil
}
