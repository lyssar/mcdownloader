package utils

import (
	"encoding/json"
	"fmt"
	"github.com/lyssar/msdcli/config"
	"github.com/manifoldco/promptui"
	"log"
)

func SelectMinecraftVersion() MinecraftVersion {
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   promptui.IconSelect + " {{ .ID | cyan }}",
		Inactive: "  {{ .ID | white }}",
		Selected: promptui.IconGood + " Mincraft Version: {{ .ID | black }}",
	}

	filter := []string{"release"}
	versionList := fetchMinecraftVersions(filter)

	prompt := promptui.Select{
		Label:     "Select Minecraft version:",
		Items:     versionList,
		Templates: templates,
	}

	i, _, err := prompt.Run()

	if err != nil {
		log.Fatal(err)
	}

	return versionList[i]
}

func GetMinecraftVersionInfo(minecraftVersionNumber string) MinecraftVersion {
	var mcVersion = MinecraftVersion{}
	mcVersions := fetchMinecraftVersions(nil)

	if len(mcVersions) > 0 {
		for _, el := range mcVersions {
			if el.ID == minecraftVersionNumber {
				return el
			}
		}
	}

	return mcVersion
}

func fetchMinecraftVersions(filter []string) []MinecraftVersion {
	var versionList []MinecraftVersion
	var dirtyList = MinecraftVersionList{}

	requestUrl := fmt.Sprintf("%s/mc/game/version_manifest_v2.json", config.MojangApiUrl)
	body := ApiCall(requestUrl)

	jsonErr := json.Unmarshal(body, &dirtyList)
	if jsonErr != nil {
		return versionList
	}

	if filter != nil {
		for _, el := range dirtyList.Versions {
			if contains(filter, el.Type) {
				versionList = append(versionList, el)
			}
		}
	} else {
		versionList = dirtyList.Versions
	}

	return versionList
}

func contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
