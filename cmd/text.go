package cmd

import (
	"github.com/spf13/cobra"
	"{{.ServiceName}}/test"
)

var (
	rootCmd = &cobra.Command{
		Use: "{{.ServiceName}}",
		PreRun: func(cmd *cobra.Command, args []string) {

		},
		Run: func(cmd *cobra.Command, args []string) {
			test.Add()
		},
	}
)
