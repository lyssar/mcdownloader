package cmd

import (
	"fmt"
	"github.com/lyssar/msdcli/curseforge/api"
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

		var selectedVersion minecraft.Version

		// when minecraft version not present
		if config.MinecraftVersion == "" {
			selectedVersion, err := minecraftMetaApi.RenderSelect()
			cobra.CheckErr(err)
			fmt.Printf("%#v", selectedVersion)
		} else {
			filter := minecraft.Filter{Minecraft: config.MinecraftVersion}
			foundVersions := minecraftMetaApi.Filter(&filter)
			if len(foundVersions) == 0 {
				selectedVersion, err := minecraftMetaApi.RenderSelect()
				cobra.CheckErr(err)
				fmt.Printf("%#v", selectedVersion)
			} else {
				selectedVersion = foundVersions[0]
			}
		}

		fmt.Printf("%#v", selectedVersion.ID)

		// when forge version no present
		//		Get Forge Version List for MC Version
		//		Make Selection
		// Download Forge server setup

		//config := utils.GetConfig()
		//curseforgeApi := curseforge.NewCurseforgeApi(getApiConfig(config))

		// versions, err := curseforgeApi.GetVersions(config.CurseForge.MinecraftGameID)
		// cobra.CheckErr(err)

		// Get filtere forge list by MC Version
		//forge, err := version.FromDefault()
		//cobra.CheckErr(err)
		//f := &version.Filter{
		//	Minecraft: "1.16.5",
		//}
		//for _, version := range forge.Filter( f) {
		//	fmt.Println(version.Minecraft)
		//}

		// fmt.Printf("%#v", forge.Releases.Filter(f))
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
