package cmd

import (
	"errors"
	"fmt"
	forgeVersionApi "github.com/kleister/go-forge/version"
	"github.com/lyssar/msdcli/curseforge/api"
	"github.com/lyssar/msdcli/forge"
	"github.com/lyssar/msdcli/minecraft"
	"github.com/lyssar/msdcli/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// forgeCmd represents the forge command
var forgeCmd = &cobra.Command{
	Use:   "forge",
	Short: "Install forge based server",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		// check args
		config := utils.GetConfig()
		minecraftMetaApi := minecraft.NewMinecraftMetaApi(config.Minecraft.MetaJson)
		minecraftMetaApi.LoadJson()

		var selectedMinecraftVersion minecraft.Version
		var err error

		if config.MinecraftVersion == "" {
			selectedMinecraftVersion, err = minecraftMetaApi.RenderSelect()
			cobra.CheckErr(err)
		} else {
			filter := minecraft.Filter{Minecraft: config.MinecraftVersion}
			foundVersions := minecraftMetaApi.Filter(&filter)
			if len(foundVersions) == 0 {
				cobra.CheckErr(
					errors.New(fmt.Sprintf("No minecraft version found for %s", config.MinecraftVersion)),
				)
			} else {
				selectedMinecraftVersion = foundVersions[0]
			}
		}

		var selectedForgeVersion *forgeVersionApi.Version

		if config.ForgeVersion == "" {
			forgeVersionList := forge.FetchForgeVersionByMinecraftVersion(selectedMinecraftVersion.ID)
			selectedForgeVersion, err = forge.RenderSelect(forgeVersionList)
		} else {
			forgeVersionList := forge.FetchForgeVersionByMinecraftAndForgeVersion(config.ForgeVersion, selectedMinecraftVersion.ID)
			if len(forgeVersionList) == 0 {
				cobra.CheckErr(
					errors.New(fmt.Sprintf("No forge package found for %s", config.ForgeVersion)),
				)
			}
			selectedForgeVersion, err = forge.RenderSelect(forgeVersionList)
			cobra.CheckErr(err)
		}

		fmt.Printf("%#v", selectedMinecraftVersion)
		fmt.Printf("%#v", selectedForgeVersion)
	},
}

func init() {
	serverCmd.AddCommand(forgeCmd)

	forgeCmd.PersistentFlags().StringP("minecraft-version", "m", "", "The minecraft version to use.")
	forgeCmd.PersistentFlags().StringP("forge-version", "f", "", "The forge version you want to use.")
	_ = viper.BindPFlag("MinecraftVersion", forgeCmd.PersistentFlags().Lookup("minecraft-version"))
	_ = viper.BindPFlag("ForgeVersion", forgeCmd.PersistentFlags().Lookup("forge-version"))
}

func getApiConfig(config *utils.Config) api.CurseforgeApiConfig {
	return api.CurseforgeApiConfig{
		BaseUrlProtocol: config.CurseForge.BaseUrlProtocol,
		BaseUrl:         config.CurseForge.BaseUrl,
		ApiKey:          config.CurseForge.ApiKey,
	}
}
