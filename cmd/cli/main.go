package main

import (
	"github.com/avocatl/admiral/pkg/commander"
	"github.com/spf13/cobra"
)

// We are just providing an example implementation to avoid
// the linters to label the base dependencies as indirects.
func main() {
	c := commander.Builder(
		nil,
		commander.Config{
			Namespace: "example",
			Execute: func(cmd *cobra.Command, args []string) {
				cmd.Println("Example command.")
			},
		},
		commander.NoCols(),
	)

	c.Execute()
}
