//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=oapi-codegen.yaml ../api.v1.yaml
package rest

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/savageking-io/necrest/kafka"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Config struct {
	Port           uint16   `yaml:"port"`
	AllowedOrigins []string `yaml:"allowed_origins"`
}

type REST struct {
	Port           uint16
	AllowedOrigins []string
	mux            *chi.Mux
}

func (d *REST) Init(config *Config) error {
	log.Traceln("REST::Init")

	log.Debugf("Config: %+v", config)
	d.Port = config.Port
	d.AllowedOrigins = config.AllowedOrigins
	if len(d.AllowedOrigins) == 0 {
		d.AllowedOrigins = []string{"http://localhost:3000"}
	}
	return nil
}

func (d *REST) Start(kafka *kafka.Kafka) error {
	log.Traceln("REST::Start")

	api := API{Kafka: kafka}

	d.mux = chi.NewMux()
	d.mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   d.AllowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	handler := HandlerFromMux(api, d.mux)

	log.Infof("Starting REST server on port %d", d.Port)
	s := &http.Server{Handler: handler, Addr: fmt.Sprintf(":%d", d.Port)}
	return s.ListenAndServe()
}
