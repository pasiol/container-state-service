package main

import (
	"net/http"
	"os"

	"github.com/pasiol/container-state-service/controllers"
)

func main() {
	controllers.RegisterControllers()
	port := ":" + os.Getenv("PORT")
	http.ListenAndServe(port, nil)
}
