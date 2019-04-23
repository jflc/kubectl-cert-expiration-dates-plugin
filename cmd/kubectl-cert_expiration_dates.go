package main

import (
	"os"

	"github.com/jflc/kubectl-cert-expiration-dates-plugin/pkg/cmd"

	"github.com/spf13/pflag"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

func main() {
	flags := pflag.NewFlagSet("kubectl-cert_expiration_dates", pflag.ExitOnError)
	pflag.CommandLine = flags

	streams := genericclioptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr}
	root := cmd.NewCmdCertExpirationDates(streams)
	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
