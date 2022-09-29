// ----------------
// Games API @see https://docs.curseforge.com/?go#curseforge-core-api-mods
// ----------------

package api

import (
	"bytes"
	"encoding/json"
	"github.com/google/go-querystring/query"
	"github.com/lyssar/msdcli/curseforge/schemas"
	"github.com/lyssar/msdcli/curseforge/utils"
	"net/http"
	"strconv"
)

const (
	UriModsSearch     schemas.ApiUri = "/v1/mods/search"
	UriMod            schemas.ApiUri = "/v1/mods/{modId}"
	UriMods           schemas.ApiUri = "/v1/mods"
	UriModsFeatured   schemas.ApiUri = "/v1/mods/featured"
	UriModDescription schemas.ApiUri = "/v1/mods/{modId}/description"
)

type SearchModsRequestData struct {
	GameID            int                         `url:"gameId"`
	ClassID           int                         `url:"classId"`
	CategoryID        int                         `url:"categoryId"`
	GameVersion       string                      `url:"gameVersion"`
	SearchFilter      string                      `url:"searchFilter"`
	SortField         schemas.ModsSearchSortField `url:"sortField"`
	SortOrder         schemas.SortOrder           `url:"sortOrder"`
	ModLoaderType     schemas.ModLoaderType       `url:"modLoaderType"`
	GameVersionTypeID int                         `url:"gameVersionTypeId"`
	Slug              string                      `url:"slug"`
	Index             int                         `url:"index"`
	PageSize          int                         `url:"pageSize"`
}

// SearchMods @see https://docs.curseforge.com/?go#search-mods
func (api CurseforgeApi) SearchMods(requestData SearchModsRequestData) (response schemas.SearchModsResponse, err error) {
	q, _ := query.Values(requestData)

	client := api.newCurseforgeClientForRoute(string(UriModsSearch))
	client.Query(&q)
	client.Request()
	err = client.GetContent(&response)

	return
}

// GetMod @see https://docs.curseforge.com/?go#get-mod
func (api CurseforgeApi) GetMod(modID int) (response schemas.GetModResponse, err error) {
	uri := utils.ReplaceNamed(string(UriMod), map[string]string{"modId": strconv.Itoa(modID)})

	client := api.newCurseforgeClientForRoute(uri)
	client.Request()
	err = client.GetContent(&response)

	return
}

// GetMods @see https://docs.curseforge.com/?go#get-mods
func (api CurseforgeApi) GetMods(requestData schemas.GetModsByIdsListRequestBody) (response schemas.GetModsResponse, err error) {
	data, err := json.Marshal(requestData)

	client := api.newCurseforgeClientForRoute(string(UriMods))
	client.Method(http.MethodPost)
	client.Data(bytes.NewReader(data))
	client.Request()
	err = client.GetContent(&response)

	return
}

// GetFeaturedMods @see https://docs.curseforge.com/?go#get-featured-mods
func (api CurseforgeApi) GetFeaturedMods(requestData schemas.GetFeaturedModsRequestBody) (response schemas.GetFeaturedModsResponse, err error) {
	data, err := json.Marshal(requestData)

	client := api.newCurseforgeClientForRoute(string(UriModsFeatured))
	client.Method(http.MethodPost)
	client.Data(bytes.NewReader(data))
	client.Request()
	err = client.GetContent(&response)
	return
}

// GetModDescription @see https://docs.curseforge.com/?go#get-mod-description
func (api CurseforgeApi) GetModDescription(modID int) (response schemas.StringResponse, err error) {
	uri := utils.ReplaceNamed(string(UriModDescription), map[string]string{"modId": strconv.Itoa(modID)})

	client := api.newCurseforgeClientForRoute(uri)
	client.Request()
	err = client.GetContent(&response)

	return
}
