package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

type logAggregator struct {
	context    string
	kubeconfig string
	namespace  string
}

func (l logAggregator) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var args []string
	if l.namespace != "" {
		args = append(args, []string{"-n", l.namespace}...)
	}
	if l.context != "" {
		args = append(args, []string{"-c", l.context}...)
	}
	if l.kubeconfig != "" {
		args = append(args, []string{"-k", l.kubeconfig}...)
	}
	runName := strings.TrimPrefix(r.URL.Path, "/")
	args = append(args, []string{"pipelinerun", "logs", runName}...)

	/*
		Make gosec exception as we call the tekton binary directly.
		Purpose of the call is to retrieve logs from tekton pipeline pods.
		No disruptive operation is possible.
		Reference issue for gosec exclusion: https://github.com/securego/gosec/issues/106
	*/
	// #nosec
	cmd := exec.CommandContext(r.Context(), "tkn", args...)
	cmd.Stdout = w
	cmd.Stderr = w
	err := cmd.Run()
	if err != nil {
		_, _ = fmt.Fprintf(w, "tkn failed: %s", err)
	}
}

func main() {
	var handler logAggregator
	flag.StringVar(&handler.context, "context", "", "name of the kubeconfig context to use (default: kubectl config current-context)")
	flag.StringVar(&handler.kubeconfig, "kubeconfig", "", "kubectl config file (default: $HOME/.kube/config)")
	flag.StringVar(&handler.namespace, "namespace", "", "namespace to use (default: from $KUBECONFIG)")
	flag.Parse()
	http.Handle("/", handler)
	s := &http.Server{Addr: ":8080"}
	log.Fatal(s.ListenAndServe())
}
