package context

import (
	"fmt"
	"io"
	"text/tabwriter"
)

// Print : write formatted context to output
func Print(output io.Writer, ctx CertExpirationDatesContext) {

	w := tabwriter.NewWriter(output, 0, 0, 3, ' ', 0)
	fmt.Fprintln(w, "CONTEXT\tCLUSTER\tUSER\tVALID FROM\tVALID TO")

	for _, c := range ctx.Certificates {
		fmt.Fprintf(w, "%v\t%v\t%v\t%v\t%v\n", c.Context, c.Cluster, c.User, c.ValidFrom, c.ValidTo)
	}
	w.Flush()
}
