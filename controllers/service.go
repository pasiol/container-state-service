package controllers

import (
	"net/http"
	"regexp"
)

type serviceController struct {
	serviceNamePattern *regexp.Regexp
}

func (sc serviceController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("container state service..."))
}

func newServiceController() *serviceController {
	return &serviceController{
		serviceNamePattern: regexp.MustCompile(`^/services/[a-z0-9]([-a-z0-9]*[a-z0-9])?/?`),
	}
}
