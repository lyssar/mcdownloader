// ----------------
// Files API @see https://docs.curseforge.com/?go#curseforge-core-api-files
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

type GetModFilesQuery struct {
	gameVersion       *string
	modLoaderType     *schemas.ModLoaderType
	gameVersionTypeID *int
	index             *int
	pageSize          *int
}

const (
	UriModFile               schemas.ApiUri = "/v1/mods/{modId}/files/{fileId}"
	UriModFiles              schemas.ApiUri = "/v1/mods/{modId}/files"
	UriGetFiles              schemas.ApiUri = "/v1/mods/files"
	UriGetModFileChangelog   schemas.ApiUri = "/v1/mods/{modId}/files/{fileId}/changelog"
	UriGetModFileDownloadURL schemas.ApiUri = "/v1/mods/{modId}/files/{fileId}/download-url"
)

// GetModFile @see https://docs.curseforge.com/?go#get-mod-file
func (api CurseforgeApi) GetModFile(modID int, fileID int) (response schemas.GetModFileResponse, err error) {
	uri := utils.ReplaceNamed(string(UriModFile), map[string]string{"modId": strconv.Itoa(modID), "fileId": strconv.Itoa(fileID)})

	client := api.newCurseforgeClientForRoute(uri)
	client.Request()
	err = client.GetContent(&response)

	return
}

// GetModFiles @see https://docs.curseforge.com/?go#get-mod-files
func (api CurseforgeApi) GetModFiles(modID int, queryData GetModFilesQuery) (response schemas.GetModFilesResponse, err error) {
	uri := utils.ReplaceNamed(string(UriModFiles), map[string]string{"modId": strconv.Itoa(modID)})
	q, _ := query.Values(queryData)

	client := api.newCurseforgeClientForRoute(uri)
	client.Query(&q)
	client.Request()
	err = client.GetContent(&response)
	return
}

// GetFiles @see https://docs.curseforge.com/?go#get-files
func (api CurseforgeApi) GetFiles(request schemas.GetModFilesRequestBody) (response schemas.GetFilesResponse, err error) {
	jsonRequest, err := json.Marshal(request)

	client := api.newCurseforgeClientForRoute(string(UriGetFiles))
	client.Method(http.MethodPost)
	client.Data(bytes.NewReader(jsonRequest))
	client.Request()
	err = client.GetContent(&response)
	return
}

// GetModFileChangelog @see https://docs.curseforge.com/?go#get-mod-file-changelog
func (api CurseforgeApi) GetModFileChangelog(modID int, fileID int) (response schemas.StringResponse, err error) {
	uri := utils.ReplaceNamed(string(UriGetModFileChangelog), map[string]string{"modId": strconv.Itoa(modID), "fileId": strconv.Itoa(fileID)})

	client := api.newCurseforgeClientForRoute(uri)
	client.Request()
	err = client.GetContent(&response)
	return
}

// GetModFileDownloadURL @see https://docs.curseforge.com/?go#get-mod-file-download-url
func (api CurseforgeApi) GetModFileDownloadURL(modID int, fileID int) (response schemas.StringResponse, err error) {
	uri := utils.ReplaceNamed(string(UriGetModFileDownloadURL), map[string]string{"modId": strconv.Itoa(modID), "fileId": strconv.Itoa(fileID)})

	client := api.newCurseforgeClientForRoute(uri)
	client.Request()
	err = client.GetContent(&response)
	return
}
