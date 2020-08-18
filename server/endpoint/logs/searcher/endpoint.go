package searcher

import (
	"bytes"
	"context"
	"net/http"
	"os/exec"
	"strings"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	kitendpoint "github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"

	"github.com/giantswarm/prow-log-aggregator/flag"
	"github.com/giantswarm/prow-log-aggregator/pkg/errors"
	"github.com/giantswarm/prow-log-aggregator/server/middleware"
	"github.com/giantswarm/prow-log-aggregator/service"
)

const (
	// Method is the HTTP method this endpoint is registered for.
	Method = "GET"
	// Name identifies the endpoint. It is aligned to the package path.
	Name = "logs"
	// Path is the HTTP request path this endpoint is registered for.
	Path = "/logs/{id}"
)

type Config struct {
	Flag       *flag.Flag
	Logger     micrologger.Logger
	Middleware *middleware.Middleware
	Service    *service.Service
	Viper      *viper.Viper
}

type Endpoint struct {
	logger     micrologger.Logger
	middleware *middleware.Middleware
	service    *service.Service

	context    string
	kubeconfig string
	namespace  string
}

func New(config Config) (*Endpoint, error) {
	if config.Logger == nil {
		return nil, microerror.Maskf(errors.InvalidConfigError, "config.Logger must not be empty")
	}
	if config.Middleware == nil {
		return nil, microerror.Maskf(errors.InvalidConfigError, "config.Middleware must not be empty")
	}

	e := &Endpoint{
		logger:     config.Logger,
		middleware: config.Middleware,
		service:    config.Service,

		context:    config.Viper.GetString(config.Flag.Kubeconfig.Context),
		kubeconfig: config.Viper.GetString(config.Flag.Kubeconfig.Kubeconfig),
		namespace:  config.Viper.GetString(config.Flag.Kubeconfig.Namespace),
	}

	return e, nil
}

func (e *Endpoint) Decoder() kithttp.DecodeRequestFunc {
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		vars := mux.Vars(r)
		id := vars["id"]

		return id, nil
	}
}

func (e *Endpoint) Encoder() kithttp.EncodeResponseFunc {
	return func(ctx context.Context, w http.ResponseWriter, response interface{}) error {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(200)
		b, ok := response.([]byte)
		if !ok {
			return microerror.Mask(errors.BadRequestError)
		}
		_, err := w.Write(b)
		return microerror.Mask(err)
	}
}

func (e *Endpoint) Endpoint() kitendpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		pipelineID, ok := request.(string)
		if !ok {
			return nil, microerror.Mask(errors.BadRequestError)
		}

		var args []string
		if e.namespace != "" {
			args = append(args, []string{"-n", e.namespace}...)
		}
		if e.context != "" {
			args = append(args, []string{"-c", e.context}...)
		}
		if e.kubeconfig != "" {
			args = append(args, []string{"-k", e.kubeconfig}...)
		}
		runName := strings.TrimPrefix(pipelineID, "/")
		args = append(args, []string{"pipelinerun", "logs", runName}...)

		var b bytes.Buffer
		/*
			Make gosec exception as we call the tekton binary directly.
			Purpose of the call is to retrieve logs from tekton pipeline pods.
			No disruptive operation is possible.
			Reference issue for gosec exclusion: https://github.com/securego/gosec/issues/106
		*/
		// #nosec
		cmd := exec.CommandContext(ctx, "tkn", args...)
		cmd.Stderr = &b
		cmd.Stdout = &b
		err := cmd.Run()
		if err != nil {
			return nil, microerror.Mask(err)
		}

		return b.String(), nil
	}
}

func (e *Endpoint) Method() string {
	return Method
}

// Middlewares returns a slice of the middlewares used in this endpoint.
func (e *Endpoint) Middlewares() []kitendpoint.Middleware {
	return []kitendpoint.Middleware{}
}

func (e *Endpoint) Name() string {
	return Name
}

func (e *Endpoint) Path() string {
	return Path
}
