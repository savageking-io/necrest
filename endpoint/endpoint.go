package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/savageking-io/necrest/service"
)

type Endpoints struct {
	PostAuthUserEndpoint endpoint.Endpoint
}

func MakeServerEndpoints(s service.Service) Endpoints {
	return Endpoints{
		PostAuthUserEndpoint: MakePostAuthUserEndpoint(s),
	}
}

func New(s service.Service) Endpoints {
	return Endpoints{
		PostAuthUserEndpoint: MakePostAuthUserEndpoint(s),
	}
}

func MakePostAuthUserEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(PostAuthUserRequest)
		v0, v1 := s.PostAuthUser(req.Username, req.Password)
		return PostAuthUserResponse{Token: v0, Err: v1}, nil
	}
}

type PostAuthUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type PostAuthUserResponse struct {
	Token string `json:"token"`
	Err   error  `json:"err,omitempty"`
}
