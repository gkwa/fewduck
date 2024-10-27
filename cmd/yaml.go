package cmd

import (
	"github.com/gkwa/fewduck/core"
	"github.com/spf13/cobra"
)

var yamlCmd = &cobra.Command{
	Use:   "yaml [path]",
	Short: "Read value of x from YAML file",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		logger := LoggerFrom(cmd.Context())
		return core.ReadYAML(logger, args[0])
	},
}

func init() {
	rootCmd.AddCommand(yamlCmd)
}
