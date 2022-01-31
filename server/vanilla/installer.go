package vanilla

import (
	"fmt"
	"github.com/lyssar/msdcli/config"
	"github.com/lyssar/msdcli/utils"
)

func DownloadInstaller() utils.MinecraftVersion {
	var minecraftVersion utils.MinecraftVersion
	if *config.McVersion == "" {
		minecraftVersion = utils.SelectMinecraftVersion()
	} else {
		minecraftVersion = utils.GetMinecraftVersionInfo(*config.McVersion)
	}

	return minecraftVersion
}

func InstalServer(minecraftVerion utils.MinecraftVersion) {
	fmt.Println("[vanilla] not implemented yet.")
}
