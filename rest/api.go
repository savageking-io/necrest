package rest

import (
	"encoding/json"
	"github.com/savageking-io/necrest/kafka"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type API struct {
	Kafka *kafka.Kafka
}

func (A API) PostAuth(w http.ResponseWriter, r *http.Request) {
	log.Traceln("API::PostAuth")

	if A.Kafka == nil {
		log.Errorf("API::PostAuth: Kafka not initialized")
		return
	}

	req := kafka.AuthRequest{
		Username: r.URL.Query().Get("username"),
		Password: r.URL.Query().Get("password"),
	}

	data, err := json.Marshal(req)
	if err != nil {
		log.Errorf("API::PostAuth: Marshal Auth Request Error: %s", err.Error())
		return
	}

	err = A.Kafka.SendMessage("auth-events", data)
	if err != nil {
		log.Errorf("API::PostAuth: Send Auth Request Error: %s", err.Error())
		return
	}
}

func (A API) GetStatus(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (A API) GetStore(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (A API) PutStore(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (A API) GetStoreStoreId(w http.ResponseWriter, r *http.Request, storeId string, params GetStoreStoreIdParams) {
	//TODO implement me
	panic("implement me")
}
