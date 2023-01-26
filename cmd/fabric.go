package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// fabricCmd represents the fabric command
var fabricCmd = &cobra.Command{
	Use:   "fabric",
	Short: "Install fabric based server",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("fabric called")
	},
}

func init() {
	setupCmd.AddCommand(fabricCmd)

	fabricCmd.Flags().String("mcversion", "", "The minecraft version to use.")
	fabricCmd.Flags().String("fabricVersion", "", "The fabric version you want to use.")
}
