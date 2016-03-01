// Package infra proxies Terraform commands.
package infra

import (
	"github.com/spf13/cobra"

	"github.com/apex/apex/cmd/apex/root"
	"github.com/apex/apex/infra"
)

var infraEnvironment string

// example output.
const example = `  View change plan
  $ apex infra plan

  Apply changes
  $ apex infra apply`

// Command config.
var Command = &cobra.Command{
	Use:     "infra",
	Short:   "Infrastructure management",
	Example: example,
	RunE:    run,
}

// Initialize.
func init() {
	root.Register(Command)
}

// Run command.
func run(c *cobra.Command, args []string) error {
	if err := root.Project.LoadFunctions(); err != nil {
		return err
	}

	p := &infra.Proxy{
		Project: root.Project,
	}

	return p.Run(args...)
}
