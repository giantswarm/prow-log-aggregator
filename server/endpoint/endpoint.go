package endpoint

import (
	"github.com/giantswarm/microendpoint/endpoint/healthz"
	"github.com/giantswarm/microendpoint/endpoint/version"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/viper"

	"github.com/giantswarm/prow-log-aggregator/flag"
	"github.com/giantswarm/prow-log-aggregator/server/endpoint/logs"
	"github.com/giantswarm/prow-log-aggregator/server/middleware"
	"github.com/giantswarm/prow-log-aggregator/service"
)

type Config struct {
	Flag       *flag.Flag
	Logger     micrologger.Logger
	Middleware *middleware.Middleware
	Service    *service.Service
	Viper      *viper.Viper
}

type Endpoint struct {
	Logs    *logs.Endpoint
	Healthz *healthz.Endpoint
	Version *version.Endpoint
}

func New(config Config) (*Endpoint, error) {
	var err error

	var logsEndpoint *logs.Endpoint
	{
		logsConfig := logs.Config{
			Flag:       config.Flag,
			Logger:     config.Logger,
			Middleware: config.Middleware,
			Viper:      config.Viper,
		}
		logsEndpoint, err = logs.New(logsConfig)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var healthzEndpoint *healthz.Endpoint
	{
		c := healthz.Config{
			Logger: config.Logger,
		}

		healthzEndpoint, err = healthz.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var versionEndpoint *version.Endpoint
	{
		versionConfig := version.Config{
			Logger:  config.Logger,
			Service: config.Service.Version,
		}

		versionEndpoint, err = version.New(versionConfig)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	e := &Endpoint{
		Logs:    logsEndpoint,
		Healthz: healthzEndpoint,
		Version: versionEndpoint,
	}

	return e, nil
}
