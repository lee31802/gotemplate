package main

import "github.com/spf13/cobra"

var (
	rootCmd = &cobra.Command{
		Use: "{{.ServiceName}}",
		PreRun: func(cmd *cobra.Command, args []string) {
			//initRoot()
		},
		Run: func(cmd *cobra.Command, args []string) {
			//internal.StartHttpServer(
			//	// add new module here ...
			//	tools.Module(),
			//)
		},
	}
)

func main() {

}
