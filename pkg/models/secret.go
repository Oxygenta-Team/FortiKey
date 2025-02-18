package models

type Secret struct {
	ID    uint64 `json:"id" db:"id"`
	Key   string `json:"key" db:"key"`
	Value string `json:"value,-"`
	Hash  []byte `json:"-" db:"hash"`
}
