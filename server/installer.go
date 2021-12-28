package server

import (
	"fmt"
	"github.com/lyssar/mcdownloader/server/forge"
	"github.com/manifoldco/promptui"
)

func InstalServer() {
	prompt := promptui.Select{
		Label: "Select server type",
		Items: []string{"Forge", "Fabric", "Spitgot", "PaperMC"},
	}
	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}
	fmt.Printf("You choose %q\n", result)

	switch result {
	case "Forge":
		minecraftVerion, forgeVersion := forge.DownloadInstaller()
		forge.InstalServer(minecraftVerion, forgeVersion)
	}
}
