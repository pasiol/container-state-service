package main

import (
	"log"
	"net/http"
	"os"

	"github.com/pasiol/container-state-service/controllers"
)

func main() {
	controllers.RegisterControllers()
	port := ":" + os.Getenv("STATE_SERVICE_PORT")
	log.Printf("Starting to listening port: %s", port)
	log.Fatal(http.ListenAndServe(port, nil))

}
