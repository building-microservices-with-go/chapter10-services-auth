package main

import (
	"net/http"
	"os"

	"github.com/building-microservices-with-go/chapter11-services-auth/handlers"
	log "github.com/sirupsen/logrus"
)

const address = ":8080"

func main() {
	var logger = &log.Logger{
		Out:       os.Stdout,
		Formatter: new(log.TextFormatter),
		Level:     log.DebugLevel,
	}

	jwt := handlers.NewJWT(logger)
	http.DefaultServeMux.HandleFunc("/", jwt.Handle)
	http.DefaultServeMux.HandleFunc("/health", handlers.HealthHandler)

	logger.WithField("service", "jwt").Infof("Starting server, listening on %s", address)
	log.WithField("service", "jwt").Fatal(http.ListenAndServe(address, http.DefaultServeMux))
}
