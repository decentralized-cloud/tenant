// Package cmd implements different commands that can be executed against project service
package cmd

import (
	"github.com/spf13/cobra"

	"github.com/micro-business/go-core/pkg/util"
)

// NewRootCommand returns root CLI application command interface
func NewRootCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use: "project",
		PreRun: func(cmd *cobra.Command, args []string) {
			printHeader()
		},
	}

	// Register all commands
	cmd.AddCommand(
		newStartCommand(),
		newVersionCommand(),
	)

	return cmd
}

func printHeader() {
	util.PrintInfo("Project Serice")
}
