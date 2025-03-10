package crypt

import "errors"

var (
	ErrBcryptCompare  = errors.New("error on bcrypt decrypting")
	ErrBcryptGenerate = errors.New("error on bcrypt encrypting")
)
