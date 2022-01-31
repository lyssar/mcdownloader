package server

import (
	"fmt"
	"strings"

	"github.com/lyssar/msdcli/config"
	"github.com/lyssar/msdcli/server/forge"
	"github.com/lyssar/msdcli/server/vanilla"
	"github.com/manifoldco/promptui"
)

func InstalServer() {
	var serverType string = *config.ServerType
	if serverType == "" {
		prompt := promptui.Select{
			Label: "Select server type",
			Items: []string{"Vanilla", "Forge", "Fabric", "Spitgot", "PaperMC"},
		}
		_, result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
		serverType = strings.ToLower(result)
		fmt.Printf("You choose %q\n", result)
	}

	switch serverType {
	case "forge":
		minecraftVersion, forgeVersion := forge.DownloadInstaller()
		forge.InstalServer(minecraftVersion, forgeVersion)
	case "vanilla":
		minecraftVersion := vanilla.DownloadInstaller()
		vanilla.InstalServer(minecraftVersion)
	}
}
