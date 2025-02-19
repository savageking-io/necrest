package service

import (
	"github.com/go-kit/log"
)

type Service interface {
	PostAuthUser(username, password string) (string, error)
}

func New(logger log.Logger) Service {
	var svc Service
	{
		svc = NewService()
	}
	return svc
}

func NewService() Service {
	return restService{}
}

type restService struct{}

func (s restService) PostAuthUser(username, password string) (string, error) {
	return "", nil
}
