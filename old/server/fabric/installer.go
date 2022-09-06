package fabric

import "fmt"

type FabricVersion struct {
	Version   string
	Installer string
}

type MinecraftVersion struct {
	Version string
	Page    string
}

func DownloadInstaller() {
	fmt.Println("[fabric] not implemented yet.")
	// TODO get version list and prompt, static list from github gist used
	// https://gist.githubusercontent.com/lyssar/430cfd38967d4e3146f8d910b1066b03/raw/

	// TODO get latest installer and install server https://meta.fabricmc.net/v2/versions/installer
	// TODO remove unnecessary stuff
}

func InstalServer() {
	fmt.Println("[fabric] not implemented yet.")
}
