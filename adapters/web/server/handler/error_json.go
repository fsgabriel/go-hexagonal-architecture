package handler

import "encoding/json"

func JsonError(msg string) []byte {
	errr := struct {
		Message string `json:"message"`
	}{
		Message: msg,
	}

	r, err := json.Marshal(errr)
	if err != nil {
		return []byte(err.Error())
	}
	return r
}
