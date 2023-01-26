package cmd

import (
	"github.com/spf13/cobra"
)

// setupCmd represents the server command
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Install a minecraft server by type. See `help` for list of types",
	Long:  "Install a minecraft server with one of the given types.",
}

func init() {
	rootCmd.AddCommand(setupCmd)
}
