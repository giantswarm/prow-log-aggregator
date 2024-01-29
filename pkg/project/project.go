package project

var (
	description string = "Extract logs from Tekton pipelines into Prow dashboards."
	gitSHA             = "n/a"
	name        string = "prow-log-aggregator"
	source      string = "https://github.com/giantswarm/prow-log-aggregator"
	version            = "0.3.2"
)

func Description() string {
	return description
}

func GitSHA() string {
	return gitSHA
}

func Name() string {
	return name
}

func Source() string {
	return source
}

func Version() string {
	return version
}
