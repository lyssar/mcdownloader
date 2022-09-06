package cmd

import (
	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Install a minecraft server by type. See `help` for list of types",
	Long:  "Install a minecraft server with one of the given types.",
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
