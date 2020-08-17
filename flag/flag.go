package flag

import (
	"github.com/giantswarm/microkit/flag"

	"github.com/giantswarm/prow-log-aggregator/flag/kubeconfig"
)

type Flag struct {
	Kubeconfig kubeconfig.Kubeconfig
}

func New() *Flag {
	f := &Flag{}
	flag.Init(f)
	return f
}
