package cmd

import (
	"fmt"
	"github.com/lyssar/msdcli/curseforge"
	"github.com/lyssar/msdcli/curseforge/api"
	"github.com/lyssar/msdcli/utils"
	"github.com/spf13/cobra"
)

// forgeCmd represents the forge command
var forgeCmd = &cobra.Command{
	Use:   "forge",
	Short: "Install forge based server",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("forge called")
		config := utils.GetConfig()
		apiConfig := api.CurseforgeApiConfig{
			BaseUrlProtocol: config.CurseForge.BaseUrlProtocol,
			BaseUrl:         config.CurseForge.BaseUrl,
			ApiKey:          config.CurseForge.ApiKey,
		}

		curseforgeApi := curseforge.NewCurseforgeApi(apiConfig)
		versions, err := curseforgeApi.GetVersions(config.CurseForge.MinecraftGameID)
		cobra.CheckErr(err)
		fmt.Println("GetGame:")
		fmt.Printf("%#v", versions)
	},
}

func init() {
	serverCmd.AddCommand(forgeCmd)

	forgeCmd.Flags().String("mcversion", "", "The minecraft version to use.")
	forgeCmd.Flags().String("forgeVersion", "", "The forge version you want to use.")
}
