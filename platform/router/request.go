package router

import (
	"encoding/json"
	"net/http"
)

func Decode(r *http.Request, val interface{}) error {
	decoder := json.NewDecoder(r.Body)

	//if you send more than expected then error
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(val); err != nil {
		return err
	}

	return nil
}
