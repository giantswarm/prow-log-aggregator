package middleware

import (
	"github.com/giantswarm/micrologger"
	"github.com/giantswarm/prow-log-aggregator/service"
)

type Config struct {
	Logger  micrologger.Logger
	Service *service.Service
}

func New(config Config) (*Middleware, error) {
	return &Middleware{}, nil
}

type Middleware struct {
}
