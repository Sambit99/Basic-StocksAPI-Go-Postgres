package utils

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(x); err != nil {
		log.Fatal("Error while parsing body", err.Error())
		return err
	}

	if decoder.More() {
		return errors.New("request body must only contain a single JSON object")
	}

	return nil

}
