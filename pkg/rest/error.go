package rest

type Error struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
