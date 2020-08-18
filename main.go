package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/microkit/command"
	microserver "github.com/giantswarm/microkit/server"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/viper"

	"github.com/giantswarm/prow-log-aggregator/flag"
	"github.com/giantswarm/prow-log-aggregator/pkg/project"
	"github.com/giantswarm/prow-log-aggregator/server"
	"github.com/giantswarm/prow-log-aggregator/service"
)

var (
	f *flag.Flag = flag.New()
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	err := mainWithError()
	if err != nil {
		panic(fmt.Sprintf("%#v\n", microerror.Mask(err)))
	}
}

func mainWithError() error {
	var err error

	// Create a new logger which is used by all packages.
	var newLogger micrologger.Logger
	{
		newLogger, err = micrologger.New(micrologger.Config{})
		if err != nil {
			return microerror.Mask(err)
		}
	}

	// We define a server factory to create the custom server once all command
	// line flags are parsed and all microservice configuration is storted out.
	newServerFactory := func(v *viper.Viper) microserver.Server {

		// Create a new custom service which implements business logic.
		var newService *service.Service
		{
			serviceConfig := service.Config{
				Description: project.Description(),
				GitCommit:   project.GitSHA(),
				Name:        project.Name(),
				Source:      project.Source(),
				Version:     project.Version(),
			}

			newService, err = service.New(serviceConfig)
			if err != nil {
				panic(fmt.Sprintf("%#v", err))
			}
		}

		// Create a new custom server which bundles our endpoints.
		var newServer microserver.Server
		{
			serverConfig := server.Config{
				Logger:  newLogger,
				Service: newService,
				Viper:   v,
				Flag:    f,

				ProjectName: project.Name(),
			}

			newServer, err = server.New(serverConfig)
			if err != nil {
				panic(fmt.Sprintf("%#v", err))
			}
		}

		return newServer
	}

	// Create a new microkit command which manages our custom microservice.
	var newCommand command.Command
	{
		c := command.Config{
			Logger:        newLogger,
			ServerFactory: newServerFactory,

			Description: project.Description(),
			GitCommit:   project.GitSHA(),
			Name:        project.Name(),
			Source:      project.Source(),
			Version:     project.Version(),
		}

		newCommand, err = command.New(c)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	daemonCommand := newCommand.DaemonCommand().CobraCommand()

	daemonCommand.PersistentFlags().String(f.Kubeconfig.Context, "", "Name of the kubeconfig context to use (default: kubectl config current-context)")
	daemonCommand.PersistentFlags().String(f.Kubeconfig.Kubeconfig, "", "Kubectl config file (default: $HOME/.kube/config).")
	daemonCommand.PersistentFlags().String(f.Kubeconfig.Namespace, "", "Namespace to use (default: from $KUBECONFIG).")

	err = newCommand.CobraCommand().Execute()
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}
