package server

import (
	"context"
	"net/http"
	"sync"

	"github.com/giantswarm/microerror"
	microserver "github.com/giantswarm/microkit/server"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/viper"

	"github.com/giantswarm/prow-log-aggregator/flag"
	"github.com/giantswarm/prow-log-aggregator/pkg/errors"
	"github.com/giantswarm/prow-log-aggregator/server/endpoint"
	"github.com/giantswarm/prow-log-aggregator/server/middleware"
	"github.com/giantswarm/prow-log-aggregator/service"
)

// Config represents the configuration used to create a new server object.
type Config struct {
	Logger  micrologger.Logger
	Service *service.Service
	Viper   *viper.Viper
	Flag    *flag.Flag

	ProjectName string
}

type Server struct {
	// Dependencies.
	logger micrologger.Logger

	// Internals.
	bootOnce     sync.Once
	config       microserver.Config
	shutdownOnce sync.Once
}

// New creates a new configured server object.
func New(config Config) (*Server, error) {
	if config.Logger == nil {
		return nil, microerror.Maskf(errors.InvalidConfigError, "%T.Logger must not be empty", config)
	}
	if config.Service == nil {
		return nil, microerror.Maskf(errors.InvalidConfigError, "%T.Service must not be empty", config)
	}
	if config.Viper == nil {
		return nil, microerror.Maskf(errors.InvalidConfigError, "%T.Viper must not be empty", config)
	}
	if config.ProjectName == "" {
		return nil, microerror.Maskf(errors.InvalidConfigError, "%T.ProjectName must not be empty", config)
	}

	var err error

	var middlewareCollection *middleware.Middleware
	{
		middlewareConfig := middleware.Config{
			Logger:  config.Logger,
			Service: config.Service,
		}
		middlewareCollection, err = middleware.New(middlewareConfig)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var endpointCollection *endpoint.Endpoint
	{
		c := endpoint.Config{
			Flag:       config.Flag,
			Logger:     config.Logger,
			Middleware: middlewareCollection,
			Service:    config.Service,
			Viper:      config.Viper,
		}

		endpointCollection, err = endpoint.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	s := &Server{
		logger: config.Logger,

		bootOnce: sync.Once{},
		config: microserver.Config{
			Logger:      config.Logger,
			ServiceName: config.ProjectName,
			Viper:       config.Viper,

			Endpoints: []microserver.Endpoint{
				endpointCollection.Logs.Searcher,
				endpointCollection.Healthz,
				endpointCollection.Version,
			},
			ErrorEncoder: errorEncoder,
		},
		shutdownOnce: sync.Once{},
	}

	return s, nil
}

func (s *Server) Boot() {
	s.bootOnce.Do(func() {
		// Here goes your custom boot logic for your server/endpoint/middleware, if
		// any.
	})
}

func (s *Server) Config() microserver.Config {
	return s.config
}

func (s *Server) Shutdown() {
	s.shutdownOnce.Do(func() {
		// Here goes your custom shutdown logic for your server/endpoint/middleware,
		// if any.
	})
}

func errorEncoder(ctx context.Context, err error, w http.ResponseWriter) {
	rErr := err.(microserver.ResponseError)
	uErr := rErr.Underlying()

	if errors.IsNotFound(uErr) {
		rErr.SetCode(microserver.CodeResourceNotFound)
		w.WriteHeader(http.StatusNotFound)

	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}
