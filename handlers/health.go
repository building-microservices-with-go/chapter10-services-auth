package handlers

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/DataDog/datadog-go/statsd"
)

type Health struct {
	logger *log.Logger
	statsd *statsd.Client
}

func (h *Health) Handle(rw http.ResponseWriter, r *http.Request) {
	h.statsd.Incr("health.success", nil, 1)

	fmt.Fprintln(rw, "OK")
}

func NewHealth(logger *log.Logger, statsd *statsd.Client) *Health {
	return &Health{
		logger: logger,
		statsd: statsd,
	}
}
