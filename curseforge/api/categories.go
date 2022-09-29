// ----------------
// Games API @see https://docs.curseforge.com/?go#curseforge-core-api-categories
// ----------------

package api

import (
	"github.com/lyssar/msdcli/curseforge/schemas"
	"net/url"
	"strconv"
)

const (
	UriCategories schemas.ApiUri = "/v1/categories"
)

// GetCategories @see https://docs.curseforge.com/?go#get-categories
func GetCategories(gameID int, classID *int) (response schemas.GetCategoriesResponse, err error) {
	q := url.Values{}
	q.Add("gameId", strconv.Itoa(gameID))
	if classID != nil {
		q.Add("classId", strconv.Itoa(*classID))
	}

	curseforgeClient := NewCurseforgeClientForRoute(string(UriCategories))
	curseforgeClient.Query(&q)
	curseforgeClient.Request()
	err = curseforgeClient.GetContent(&response)

	return
}
