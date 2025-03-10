package models

type SecretMethod string

var (
	BCRYPT SecretMethod = "bcrypt"
)

type Secret struct {
	ID     uint64       `json:"id" db:"id"`
	UserID uint64       `json:"user_id" db:"user_id"`
	Key    string       `json:"key" db:"key"`
	Value  string       `json:"value,-"`
	Method SecretMethod `json:"method" db:"method"`
	Hash   []byte       `json:"-" db:"hash"`
}

type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
