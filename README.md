<!--

    TODO:

    - Change the badge (with style=shield):
      https://circleci.com/gh/giantswarm/REPOSITORY_NAME/edit#badges
      If this is a private repository token with scope `status` will be needed.

    - Update CODEOWNERS file according to the needs for this project

    - Run `devctl replace -i "REPOSITORY_NAME" "$(basename $(git rev-parse --show-toplevel))" *.md`
      and commit your changes.

    - If the repository is public consider adding godoc badge. This should be
      the first badge separated with a single space.
      [![GoDoc](https://godoc.org/github.com/giantswarm/REPOSITORY_NAME?status.svg)](http://godoc.org/github.com/giantswarm/REPOSITORY_NAME)

-->
[![CircleCI](https://circleci.com/gh/giantswarm/prow-log-aggregator.svg?style=shield)](https://circleci.com/gh/giantswarm/prow-log-aggregator)

# prow-log-aggregator

This is a small tool to extract logs from our Tekton pipeline into Prow dashboards.
