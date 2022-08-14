package cmd

import (
	"github.com/spf13/cobra"
	"spin/file"
)

var (
	debugFile int
)

func debugCmd() *cobra.Command {

	command := &cobra.Command{
		Use:    "debug [-file n]",
		Short:  "debug tailspin",
		Hidden: true,
		Run: func(cmd *cobra.Command, args []string) {
			config := getConfig()
			config.DebugMode = true

			file.Setup(config)
		},
	}

	command.PersistentFlags().IntVarP(&debugFile, "file", "e", 0,
		"select a specific log file for debugging")

	return command
}
