[![CircleCI](https://circleci.com/gh/giantswarm/prow-log-aggregator.svg?style=shield)](https://circleci.com/gh/giantswarm/prow-log-aggregator)

# prow-log-aggregator

This is a small tool to extract logs from our Tekton pipeline into [Prow dashboards](https://prow.giantswarm.io/).

It wraps the Tekton CLI command line tool under the `logs` endpoint to fetch logs from Tekton pipeline runs.

## Related repositories

- https://github.com/giantswarm/test-infra/
