package cmd

import (
	"github.com/lyssar/msdcli/config"
	"github.com/lyssar/msdcli/curseforge/api"
	"github.com/lyssar/msdcli/forge"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

// forgeCmd represents the forge command
var forgeCmd = &cobra.Command{
	Use:   "forge",
	Short: "Install forge based server",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		// check args
		config := config.GetConfig()
		dir := viper.GetString("working-dir")
		if dir == "" {
			dir, _ = os.Getwd()
		}

		forge.CreateServer(dir, config)
	},
}

func init() {
	setupCmd.AddCommand(forgeCmd)

	forgeCmd.PersistentFlags().StringP("minecraft-version", "m", "", "The minecraft version to use.")
	forgeCmd.PersistentFlags().StringP("forge-version", "f", "", "The forge version you want to use.")
	_ = viper.BindPFlag("MinecraftVersion", forgeCmd.PersistentFlags().Lookup("minecraft-version"))
	_ = viper.BindPFlag("ForgeVersion", forgeCmd.PersistentFlags().Lookup("forge-version"))
}

func getApiConfig(config *config.Config) api.CurseforgeApiConfig {
	return api.CurseforgeApiConfig{
		BaseUrlProtocol: config.CurseForge.BaseUrlProtocol,
		BaseUrl:         config.CurseForge.BaseUrl,
		ApiKey:          config.CurseForge.ApiKey,
	}
}
