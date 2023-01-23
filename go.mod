module github.com/giantswarm/prow-log-aggregator

go 1.14

require (
	github.com/giantswarm/microendpoint v1.0.0
	github.com/giantswarm/microerror v0.4.0
	github.com/giantswarm/microkit v1.0.0
	github.com/giantswarm/micrologger v1.0.0
	github.com/go-kit/kit v0.12.0
	github.com/gorilla/mux v1.8.0
	github.com/pelletier/go-toml v1.9.5 // indirect
	github.com/spf13/viper v1.15.0
)

replace (
	github.com/coreos/etcd v3.3.10+incompatible => github.com/coreos/etcd v3.3.25+incompatible
	github.com/coreos/etcd v3.3.13+incompatible => github.com/coreos/etcd v3.3.25+incompatible
	github.com/dgrijalva/jwt-go => github.com/form3tech-oss/jwt-go v3.2.1+incompatible
	github.com/nats-io/nats-server/v2 => github.com/nats-io/nats-server/v2 v2.9.11
)
