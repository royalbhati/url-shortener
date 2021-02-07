package router

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

func Response(w http.ResponseWriter, val interface{}, statusCode int) error {

	//special case when we areupdating
	if statusCode == http.StatusNoContent {
		w.WriteHeader(statusCode)
		return nil
	}

	data, err := json.Marshal(val)
	if err != nil {
		return errors.Wrap(err, "Marshalling")
	}

	//The order matters for the two lines below
	w.Header().Set("content-type", "application/json; charset=utf-8")
	//this lines writes out the header so if we do anything to the headers
	//after this line it wont matter
	w.WriteHeader(statusCode)
	if _, err := w.Write(data); err != nil {
		return errors.Wrap(err, "Writing to client")
	}
	return nil
}
