package curseforge

import "github.com/lyssar/msdcli/curseforge/api"

func NewCurseforgeApi(config api.CurseforgeApiConfig) api.CurseforgeApi {
	cfApi := api.CurseforgeApi{Config: config}

	return cfApi
}
