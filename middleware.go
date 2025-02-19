package main

/*
import (
	"time"

	"github.com/go-kit/log"
	"github.com/savageking-io/necrest/service"
)

type ServiceMiddleware func(service.Service) service.Service

func LoggingMiddleware(logger log.Logger) ServiceMiddleware {
	return func(next service.Service) service.Service {
		return &loggingMiddleware{next: next, logger: logger}
	}
}

type loggingMiddleware struct {
	next   service.Service
	logger log.Logger
}

func (mw loggingMiddleware) PostAuthUser(username, password string) (username, password string) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "PostProfile", "id", p.ID, "took", time.Since(begin), "err", err)
	}(time.Now())
	return mw.next.PostAuthUser(username, password)
}*/
