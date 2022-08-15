FROM alpine:3.16.2

RUN apk add --no-cache ca-certificates curl tar

ADD prow-log-aggregator /prow-log-aggregator

ARG tkn_version=0.11.0
RUN export URL=https://github.com/tektoncd/cli/releases/download/v${tkn_version}/tkn_${tkn_version}_Linux_x86_64.tar.gz && \
    curl -sSL $URL | tar xzvf - -C /usr/local/bin tkn

EXPOSE 8000
ENTRYPOINT ["/prow-log-aggregator"]
