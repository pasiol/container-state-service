package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

func RegisterControllers() {
	sc := newServiceController()

	http.Handle("/services", *sc)
	http.Handle("/services/", *sc)
}

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
