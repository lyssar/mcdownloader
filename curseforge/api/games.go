// ----------------
// Games API @see https://docs.curseforge.com/?php#curseforge-core-api-games
// ----------------

package api

import (
	"github.com/lyssar/msdcli/curseforge"
	"github.com/lyssar/msdcli/curseforge/schemas"
	"net/url"
	"strconv"
)

const (
	UriGame              schemas.ApiUri = "/v1/games/{gameId}"
	UriGames             schemas.ApiUri = "/v1/games"
	UriGamesVersions     schemas.ApiUri = "/v1/games/{gameId}/versions"
	UriGamesVersionTypes schemas.ApiUri = "/v1/games/{gameId}/version-types"
)

func GetGames(index int, pageSize int) (response schemas.GetGamesResponse, err error) {
	q := url.Values{}
	q.Add("index", strconv.Itoa(index))
	q.Add("pageSize", strconv.Itoa(pageSize))

	client := NewCurseforgeClientForRoute(string(UriGames))
	client.Query(&q)
	client.Request()
	err = client.GetContent(&response)

	return
}

func GetGame(gameId int) (game schemas.Game, err error) {
	var responseData schemas.GetGameResponse
	uri := curseforge.ReplaceNamed(string(UriGame), map[string]string{"gameId": strconv.Itoa(gameId)})

	client := NewCurseforgeClientForRoute(uri)
	client.Request()
	err = client.GetContent(&responseData)
	game = responseData.Data

	return
}

func GetVersions(gameId int) (response schemas.GetVersionsResponse, err error) {
	uri := curseforge.ReplaceNamed(string(UriGamesVersions), map[string]string{"gameId": strconv.Itoa(gameId)})
	client := NewCurseforgeClientForRoute(uri)
	client.Request()
	err = client.GetContent(&response)
	return
}

func GetVersionTypes(gameId int) (response schemas.GetVersionTypesResponse, err error) {
	uri := curseforge.ReplaceNamed(string(UriGamesVersionTypes), map[string]string{"gameId": strconv.Itoa(gameId)})
	client := NewCurseforgeClientForRoute(uri)
	client.Request()
	err = client.GetContent(&response)
	return
}
