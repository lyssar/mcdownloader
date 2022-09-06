package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// ModpackCmd represents the modpack command
var modpackCmd = &cobra.Command{
	Use:   "modpack",
	Short: "Download and unpack a modpack from curseforge",
	Long:  "Download and unpack a modpack from curseforge",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("modpack called")
	},
}

func init() {
	rootCmd.AddCommand(modpackCmd)

	modpackCmd.Flags().Int("packageId", 0, "Modpack ID from curseforge")
	modpackCmd.Flags().Int("serverPackageFileID", 0, "File ID to download (excplicit server version of a package)")
}
