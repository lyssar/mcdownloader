package forge

import (
	"fmt"
	forgeVersionApi "github.com/kleister/go-forge/version"
	"github.com/lyssar/msdcli/minecraft"
	"github.com/lyssar/msdcli/utils"
	"github.com/spf13/cobra"
)

func CreateServer(config *utils.Config) {
	minecraftMetaApi := minecraft.NewMinecraftMetaApi(config.Minecraft.MetaJson)
	selectedMinecraftVersion, err := minecraftMetaApi.FindMinecraftVersion(config.MinecraftVersion)
	cobra.CheckErr(err)

	forgeApp := NewForgeClient()

	var selectedForgeVersion *forgeVersionApi.Version
	selectedForgeVersion, err = forgeApp.SelectForgeVersion(config.ForgeVersion, selectedMinecraftVersion.ID)
	cobra.CheckErr(err)

	fmt.Printf("%#v", selectedMinecraftVersion)
	fmt.Printf("%#v", selectedForgeVersion)
}
