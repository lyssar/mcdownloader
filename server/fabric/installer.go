package fabric

import (
	"encoding/json"
	"fmt"
	"github.com/lyssar/msdcli/config"
	"github.com/lyssar/msdcli/utils"
	"log"
)

func DownloadInstaller() {
	fmt.Println("[fabric] loading not implemented yet.")

	var minecraftVersion utils.MinecraftVersion
	if *config.McVersion == "" {
		minecraftVersion = utils.SelectMinecraftVersion()
	} else {
		minecraftVersion = utils.GetMinecraftVersionInfo(*config.McVersion)
	}

	var loader utils.FabricLoader
	if *config.ServerVersion == "" {
		loaderList := fetchStableLoaderList(minecraftVersion)
		loader = selectServerVersionFromList(loaderList)
	} else {
		loader = fetchSpecificServerVersion(config.ServerVersion)
	}

	fmt.Printf("%+v\n", loader)
	fmt.Printf("%+v\n", minecraftVersion)

	donwloadFabricFiles(minecraftVersion, loader)
}

func InstalServer() {
	fmt.Println("[fabric] install not implemented yet.")
}

func fetchStableLoaderList(minecraftVersion utils.MinecraftVersion) utils.FabricLoaderList {
	var loaderList utils.FabricLoaderList
	var dirtyLoader utils.FabricLoaderList
	body := utils.FabricApiCall(fmt.Sprintf("versions/loader/%s", minecraftVersion.ID), "v1")

	jsonErr := json.Unmarshal(body, &dirtyLoader)
	if jsonErr != nil {
		return loaderList
	}

	for _, el := range dirtyLoader {
		if el.Loader.Stable == true {
			loaderList = append(loaderList, el)
		}
	}

	return loaderList
}

func selectServerVersionFromList(loaderList utils.FabricLoaderList) utils.FabricLoader {
	var chosenLoader utils.FabricLoader

	return chosenLoader
}

func fetchSpecificServerVersion(serverVersion *string) utils.FabricLoader {
	var loader utils.FabricLoader

	return loader
}

func fetchLatestInstallerVersion() utils.FabricInstaller {
	var installerList utils.FabricInstallerList

	body := utils.FabricApiCall("versions/installer", "v2")

	jsonErr := json.Unmarshal(body, &installerList)
	if jsonErr != nil {
		log.Fatal(jsonErr.Error())
	}

	return installerList[0]
}

func donwloadFabricFiles(minecraftVersion utils.MinecraftVersion, loader utils.FabricLoader) {
	latestInstaller := fetchLatestInstallerVersion()
	downloadUrl := fmt.Sprintf("%s/versions/loader/%s/%s/%s/server/jar", config.FabricApiV2BaseUrl, minecraftVersion.ID, loader.Version, latestInstaller.Version)

	fmt.Println("Downloading fabric version: ", loader.Version)
	utils.DownloadInstaller(downloadUrl)
}
