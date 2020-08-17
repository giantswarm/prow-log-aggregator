package logs

import (
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/viper"

	"github.com/giantswarm/prow-log-aggregator/flag"
	"github.com/giantswarm/prow-log-aggregator/server/endpoint/logs/searcher"
	"github.com/giantswarm/prow-log-aggregator/server/middleware"
)

type Config struct {
	Flag       *flag.Flag
	Logger     micrologger.Logger
	Middleware *middleware.Middleware
	Viper      *viper.Viper
}

type Endpoint struct {
	Searcher *searcher.Endpoint
}

func New(config Config) (*Endpoint, error) {
	var err error

	var searcherEndpoint *searcher.Endpoint
	{
		searcherConfig := searcher.Config{
			Flag:       config.Flag,
			Logger:     config.Logger,
			Middleware: config.Middleware,
			Viper:      config.Viper,
		}
		searcherEndpoint, err = searcher.New(searcherConfig)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	newEndpoint := &Endpoint{
		Searcher: searcherEndpoint,
	}

	return newEndpoint, nil
}
