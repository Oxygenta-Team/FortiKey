package models

import "encoding/json"

const (
	CipherObjectType = "Cipher"
)

const (
	CreateActionType = "Create"
)

type KafkaMessage struct {
	ID         uint64           `json:"id"`
	ObjectType string           `json:"objectType"`
	ActionType string           `json:"actionType"`
	Object     *json.RawMessage `json:"object"`
}
