package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"

	"github.com/pasiol/container-state-service/models"
)

type serviceController struct {
	serviceNamePattern *regexp.Regexp
}

func (sc serviceController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Print("http listening.")
	if r.URL.Path == "/services" {
		if r.Method == http.MethodPost {
			sc.post(w, r)
		}
	} else {
		matches := sc.serviceNamePattern.FindStringSubmatch(r.URL.Path)
		if len(matches) == 0 {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		name := matches[1]
		switch r.Method {
		case http.MethodGet:
			sc.get(name, w)
		case http.MethodPut:
			sc.put(w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	}
}

func (sc *serviceController) put(w http.ResponseWriter, r *http.Request) {
	s, err := sc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse Service object"))
		return
	}
	s, err = models.SetEnded(s.Name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse Service object"))
		return
	}
	encodeResponseAsJSON(s, w)
}

func (sc *serviceController) get(name string, w http.ResponseWriter) {
	s, err := models.GetServiceByName(name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	encodeResponseAsJSON(s, w)
}

func (sc *serviceController) post(w http.ResponseWriter, r *http.Request) {
	s, err := sc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse Service object"))
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
		serviceNamePattern: regexp.MustCompile(`^/services/([-a-z0-9]*[a-z0-9])?/?`),
	}
}
