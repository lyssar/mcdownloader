package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// vanillaCmd represents the vanilla command
var vanillaCmd = &cobra.Command{
	Use:   "vanilla",
	Short: "Install a vanilla server",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("vanilla called")
	},
}

func init() {
	serverCmd.AddCommand(vanillaCmd)
	vanillaCmd.Flags().String("mcversion", "", "The minecraft version to use.")
}
