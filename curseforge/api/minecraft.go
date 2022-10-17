// ----------------
// Minecraft API @see https://docs.curseforge.com/?go#curseforge-core-api-fingerprints
// ----------------

package api

import (
	"github.com/lyssar/msdcli/curseforge/schemas"
	"github.com/lyssar/msdcli/curseforge/utils"
	"net/url"
)

const (
	UriGetMinecraftVersions          schemas.ApiUri = "/v1/minecraft/version"
	UriGetSpecificMinecraftVersion   schemas.ApiUri = "/v1/minecraft/version/{gameVersionString}"
	UriGetMinecraftModLoaders        schemas.ApiUri = "/v1/minecraft/modloader"
	UriGetSpecificMinecraftModLoader schemas.ApiUri = "/v1/minecraft/modloader/{modLoaderName}"
)

// GetMinecraftVersions @see https://docs.curseforge.com/?go#get-minecraft-versions
func (api CurseforgeApi) GetMinecraftVersions(sortDescending *string) (response schemas.ApiResponseOfListOfMinecraftGameVersion, err error) {
	q := url.Values{}

	if sortDescending != nil {
		q.Add("sortDescending", *sortDescending)
	}

	curseforgeClient := api.newCurseforgeClientForRoute(string(UriGetMinecraftVersions))
	curseforgeClient.Query(&q)
	curseforgeClient.Request()
	err = curseforgeClient.GetContent(&response)
	return
}

// GetSpecificMinecraftVersion @see https://docs.curseforge.com/?go#get-specific-minecraft-version
func (api CurseforgeApi) GetSpecificMinecraftVersion(gameVersionString string) (response schemas.ApiResponseOfMinecraftGameVersion, err error) {
	uri := utils.ReplaceNamed(string(UriGetSpecificMinecraftVersion), map[string]string{"gameVersionString": gameVersionString})

	client := api.newCurseforgeClientForRoute(uri)
	client.Request()
	err = client.GetContent(&response)

	return
}

// GetMinecraftModLoaders @see https://docs.curseforge.com/?go#get-minecraft-modloaders
func (api CurseforgeApi) GetMinecraftModLoaders(version *string, includeAll *string) (response schemas.ApiResponseOfListOfMinecraftModLoaderIndex, err error) {
	q := url.Values{}

	if version != nil {
		q.Add("version", *version)
	}

	if includeAll != nil {
		q.Add("includeAll", *includeAll)
	}

	curseforgeClient := api.newCurseforgeClientForRoute(string(UriGetMinecraftModLoaders))
	curseforgeClient.Query(&q)
	curseforgeClient.Request()
	err = curseforgeClient.GetContent(&response)

	return
}

// GetSpecificMinecraftModLoader @see https://docs.curseforge.com/?go#get-specific-minecraft-modloader
func (api CurseforgeApi) GetSpecificMinecraftModLoader(modLoaderName string) (response schemas.ApiResponseOfMinecraftModLoaderVersion, err error) {
	uri := utils.ReplaceNamed(string(UriGetSpecificMinecraftModLoader), map[string]string{"modLoaderName": modLoaderName})

	client := api.newCurseforgeClientForRoute(uri)
	client.Request()
	err = client.GetContent(&response)
	return
}
