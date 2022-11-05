package forge

import (
	forgeVersionApi "github.com/kleister/go-forge/version"
	"github.com/spf13/cobra"
	"sort"
)

func FetchForgeVersionByMinecraftVersion(minecraftVersionString string) forgeVersionApi.Versions {
	forge, err := forgeVersionApi.FromDefault()
	cobra.CheckErr(err)

	forgeMinecraftVersionFilter := &forgeVersionApi.Filter{
		Minecraft: minecraftVersionString,
	}

	forgeVersionList := forge.Releases.Filter(forgeMinecraftVersionFilter)

	return reverseSortVersionList(forgeVersionList)
}

func FetchForgeVersionByMinecraftAndForgeVersion(forgeVersionString string, minecraftVersionString string) forgeVersionApi.Versions {
	forge, err := forgeVersionApi.FromDefault()
	cobra.CheckErr(err)

	forgeMinecraftVersionFilter := &forgeVersionApi.Filter{
		Version:   forgeVersionString,
		Minecraft: minecraftVersionString,
	}

	forgeVersionList := forge.Releases.Filter(forgeMinecraftVersionFilter)

	return reverseSortVersionList(forgeVersionList)
}

func reverseSortVersionList(forgeVersionList forgeVersionApi.Versions) forgeVersionApi.Versions {
	sort.Slice(forgeVersionList, func(i, j int) bool {
		return forgeVersionList[i].ID > forgeVersionList[j].ID
	})

	return forgeVersionList
}
