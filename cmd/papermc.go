package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// papermcCmd represents the papermc command
var papermcCmd = &cobra.Command{
	Use:   "papermc",
	Short: "Install papermc based server",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("papermc called")
	},
}

func init() {
	setupCmd.AddCommand(papermcCmd)
	papermcCmd.Flags().String("mcversion", "", "The minecraft version to use.")
	papermcCmd.Flags().String("papermcVersion", "", "The papermc version you want to use.")
}
