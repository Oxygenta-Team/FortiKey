package testassets

import "encoding/json"

func Marshal(dest any) (b []byte) {
	b, _ = json.Marshal(dest)
	return
}
