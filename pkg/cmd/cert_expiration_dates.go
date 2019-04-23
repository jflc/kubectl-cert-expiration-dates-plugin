package cmd

import (
	"fmt"

	"github.com/jflc/kubectl-cert-expiration-dates-plugin/pkg/context"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

const cmdExample = "%[1]s cert-expiration-dates"

// NewCmdCertExpirationDates : provides a cobra command wrapping CertExpirationDatesOptions
func NewCmdCertExpirationDates(streams genericclioptions.IOStreams) *cobra.Command {
	configFlags := genericclioptions.NewConfigFlags(false)

	cmd := &cobra.Command{
		Use:          "cert-expiration-dates [flags]",
		Short:        "View certificate expiration dates",
		Example:      fmt.Sprintf(cmdExample, "kubectl"),
		SilenceUsage: true,
		RunE: func(c *cobra.Command, args []string) error {
			builder := context.Builder()
			ctx, err := builder.ConfigFlags(configFlags).Build()
			if err != nil {
				return err
			}

			context.Print(streams.Out, ctx)

			return nil
		},
	}

	configFlags.AddFlags(cmd.Flags())

	return cmd
}
