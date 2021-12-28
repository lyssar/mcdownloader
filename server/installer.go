package server

import (
	"fmt"
	"github.com/lyssar/mcdownloader/config"
	"github.com/lyssar/mcdownloader/server/forge"
	"github.com/manifoldco/promptui"
	"strings"
)

func InstalServer() {
	var serverType string = *config.ServerType
	if serverType == "" {
		prompt := promptui.Select{
			Label: "Select server type",
			Items: []string{"Forge", "Fabric", "Spitgot", "PaperMC"},
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
		minecraftVerion, forgeVersion := forge.DownloadInstaller()
		forge.InstalServer(minecraftVerion, forgeVersion)
	}
}
