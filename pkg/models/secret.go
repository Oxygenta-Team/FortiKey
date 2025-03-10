package models

type Secret struct {
	ID   uint64 `json:"id" db:"id"`
	Key  string `json:"key" db:"key"`
	Hash []byte `json:"hash" db:"hash"`
}
