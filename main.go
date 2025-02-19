package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/go-kit/log"
	"github.com/oklog/run"
	"github.com/savageking-io/necrest/service"
	"github.com/savageking-io/necrest/transport.go"
)

func main() {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	var s service.Service
	{
		s = service.New(logger)
		//s = service.LoggingMiddleware(logger)
	}

	var h http.Handler
	{
		h = transport.NewHTTPHandler(s, logger)
	}

	var g run.Group
	g.Add(func() error {
		logger.Log("msg", "HTTP", "addr", ":8080")
		return http.ListenAndServe(":8080", h)
	}, func(error) {
		logger.Log("msg", "HTTP", "addr", ":8080", "status", "down")
	})

	interrupt := make(chan struct{})
	g.Add(func() error {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		select {
		case sig := <-c:
			return fmt.Errorf("received signal %s", sig)
		case <-interrupt:
			return nil
		}
	}, func(error) {
		close(interrupt)
	})

	g.Run()
}
