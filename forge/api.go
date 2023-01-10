package forge

import (
	"errors"
	"fmt"
	forgeVersionApi "github.com/kleister/go-forge/version"
	"github.com/lyssar/msdcli/utils"
	"github.com/spf13/cobra"
	"sort"
)

type ForgeApi struct {
	forgeVersionList forgeVersionApi.Response
	ForgeVersion     string
	MinecraftVersion string
}

func NewForgeClient() ForgeApi {
	forgeResultSet, err := forgeVersionApi.FromDefault()
	cobra.CheckErr(err)
	return ForgeApi{
		forgeVersionList: forgeResultSet,
	}
}

func (api ForgeApi) VerifyForgeVersion(forgeVersion forgeVersionApi.Version, minecraftVersionString string) (*bool, error) {
	ok := true

	if forgeVersion.Minecraft != minecraftVersionString {
		err := errors.New(
			fmt.Sprintf(
				"Forge version %s is not for minecraft %s. Expecting minecraft version %s",
				forgeVersion.ID,
				minecraftVersionString,
				forgeVersion.Minecraft,
			),
		)

		return nil, err
	}

	return &ok, nil
}

func (api ForgeApi) GetFilteredReleasesForMinecraftVersion(minecraftVersionString string) forgeVersionApi.Versions {
	forgeMinecraftVersionFilter := &forgeVersionApi.Filter{
		Minecraft: minecraftVersionString,
	}

	forgeVersionList := api.forgeVersionList.Releases.Filter(forgeMinecraftVersionFilter)

	return api.reverseSortVersionList(forgeVersionList)
}

func (api ForgeApi) GetSpecificForgeVersion(forgeVersionString string) forgeVersionApi.Versions {
	forgeMinecraftVersionFilter := &forgeVersionApi.Filter{
		Version: forgeVersionString,
	}

	forgeVersionList := api.forgeVersionList.Releases.Filter(forgeMinecraftVersionFilter)

	return api.reverseSortVersionList(forgeVersionList)
}

func (api ForgeApi) reverseSortVersionList(forgeVersionList forgeVersionApi.Versions) forgeVersionApi.Versions {
	sort.Slice(forgeVersionList, func(i, j int) bool {
		return forgeVersionList[i].ID > forgeVersionList[j].ID
	})

	return forgeVersionList
}

func (api ForgeApi) SelectForgeVersion(forgeVersionString string, minecraftVersionString string) (*forgeVersionApi.Version, error) {
	utils.PrintInfo("Choosing forge version")
	var selectedForgeVersion *forgeVersionApi.Version
	var err error

	if forgeVersionString == "" {
		forgeVersionList := api.GetFilteredReleasesForMinecraftVersion(minecraftVersionString)
		selectedForgeVersion, err = RenderSelect(forgeVersionList)
		if err != nil {
			return nil, err
		}
	} else {
		forgeVersionList := api.GetSpecificForgeVersion(forgeVersionString)

		if len(forgeVersionList) == 0 {
			forgeError := errors.New(fmt.Sprintf("No forge package found for %s", forgeVersionString))
			return nil, forgeError
		}
		_, err := api.VerifyForgeVersion(forgeVersionList[0], minecraftVersionString)

		if err != nil {
			return nil, err
		}
		selectedForgeVersion = &forgeVersionList[0]

	}

	return selectedForgeVersion, nil
}
