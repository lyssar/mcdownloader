package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// spigotCmd represents the spigot command
var spigotCmd = &cobra.Command{
	Use:   "spigot",
	Short: "Install spigot based server",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("spigot called")
	},
}

func init() {
	setupCmd.AddCommand(spigotCmd)
	spigotCmd.Flags().String("mcversion", "", "The minecraft version to use.")
	spigotCmd.Flags().String("spigotVersion", "", "The spigot version you want to use.")
}
