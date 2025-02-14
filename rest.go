//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=oapi-codegen.yaml api.v1.yaml

package main

import log "github.com/sirupsen/logrus"

type REST struct {
	Port uint16
}

func (d *REST) Init(config *RESTConfig) error {
	log.Traceln("REST::Init")
	d.Port = config.Port
	return nil
}
