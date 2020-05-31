package main

import (
	"net/http"

	"github.com/pasiol/container-state-service/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
