package service

import (
	"sync"
	"time"

	"github.com/giantswarm/microendpoint/service/version"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
)

const (
	DefaultRetryCount = 5
	DefaultTimeout    = 5 * time.Second
)

type Config struct {
	Logger micrologger.Logger

	Description string
	GitCommit   string
	Name        string
	Source      string
	Version     string
}

type Service struct {
	Version *version.Service

	bootOnce sync.Once
}

func New(config Config) (*Service, error) {
	var err error

	var versionService *version.Service
	{
		versionConfig := version.Config{
			Description: config.Description,
			GitCommit:   config.GitCommit,
			Name:        config.Name,
			Source:      config.Source,
			Version:     config.Version,
		}

		versionService, err = version.New(versionConfig)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	s := &Service{
		Version: versionService,

		bootOnce: sync.Once{},
	}

	return s, nil
}
