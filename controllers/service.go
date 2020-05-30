package controllers

import (
	"encoding/json"
	"net/http"
	"regexp"

	"../models"
)

type serviceController struct {
	serviceNamePattern *regexp.Regexp
}

func (sc serviceController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/services" {
		switch r.Method {
		case http.MethodPost:
			sc.post(w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	}
}

func (sc *serviceController) post(w http.ResponseWriter, r *http.Request) {
	s, err := sc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse User object"))
		return
	}
	s, err = models.AddService(s)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsJSON(s, w)
}

func (sc *serviceController) parseRequest(r *http.Request) (models.Service, error) {
	dec := json.NewDecoder(r.Body)
	var s models.Service
	err := dec.Decode(&s)
	if err != nil {
		return models.Service{}, err
	}
	return s, nil
}

func newServiceController() *serviceController {
	return &serviceController{
		serviceNamePattern: regexp.MustCompile(`^/services/[a-z0-9]([-a-z0-9]*[a-z0-9])?/?`),
	}
}
