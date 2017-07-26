package main

import (
	"net/http"
	"os"

	"github.com/DataDog/datadog-go/statsd"
	"github.com/building-microservices-with-go/chapter10-services-auth/handlers"
	log "github.com/sirupsen/logrus"
)

const address = ":8080"

func main() {
	var logger = &log.Logger{
		Out:       os.Stdout,
		Formatter: new(log.TextFormatter),
		Level:     log.DebugLevel,
	}

	c, err := statsd.New("127.0.0.1:8125")
	if err != nil {
		log.Fatal(err)
	}
	// prefix every metric with the app name
	c.Namespace = "chapter10.auth."

	jwt := handlers.NewJWT(logger, c)
	health := handlers.NewHealth(logger, c)

	http.DefaultServeMux.HandleFunc("/", jwt.Handle)
	http.DefaultServeMux.HandleFunc("/health", health.Handle)

	logger.WithField("service", "jwt").Infof("Starting server, listening on %s", address)
	log.WithField("service", "jwt").Fatal(http.ListenAndServe(address, http.DefaultServeMux))
}
