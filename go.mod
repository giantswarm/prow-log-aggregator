module github.com/giantswarm/prow-log-aggregator

go 1.14

require (
	github.com/giantswarm/microendpoint v0.2.0
	github.com/giantswarm/microerror v0.3.0
	github.com/giantswarm/microkit v0.2.2
	github.com/giantswarm/micrologger v0.5.0
	github.com/go-kit/kit v0.11.0
	github.com/gorilla/mux v1.8.0
	github.com/spf13/viper v1.8.1
)

replace (
	github.com/coreos/etcd v3.3.10+incompatible => github.com/coreos/etcd v3.3.25+incompatible
	github.com/coreos/etcd v3.3.13+incompatible => github.com/coreos/etcd v3.3.25+incompatible
	github.com/gogo/protobuf v1.2.1 => github.com/gogo/protobuf v1.3.2
)
