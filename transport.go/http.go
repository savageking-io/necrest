package transport

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/log"

	"github.com/go-kit/kit/transport"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/savageking-io/necrest/endpoint"
	"github.com/savageking-io/necrest/service"
)

func NewHTTPHandler(s service.Service, logger log.Logger) http.Handler {
	r := mux.NewRouter()
	e := endpoint.MakeServerEndpoints(s)

	options := []httptransport.ServerOption{
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		httptransport.ServerErrorEncoder(encodeError),
	}

	r.Methods("POST").Path("/auth").Handler(httptransport.NewServer(
		e.PostAuthUserEndpoint,
		decodePostAuthUserRequest,
		encodeResponse,
		options...,
	))
	return r
}

func decodePostAuthUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.PostAuthUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
