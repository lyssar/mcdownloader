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
func (api CurseforgeApi) GetCategories(gameID int, classID *int, classesOnly *bool) (response schemas.GetCategoriesResponse, err error) {
	q := url.Values{}
	q.Add("gameId", strconv.Itoa(gameID))
	if classID != nil {
		q.Add("classId", strconv.Itoa(*classID))
	}

	if classesOnly != nil {
		classesOnlyParam := "false"
		if *classesOnly == true {
			classesOnlyParam = "true"
		}
		q.Add("classesOnly", classesOnlyParam)
	}

	curseforgeClient := api.newCurseforgeClientForRoute(string(UriCategories))
	curseforgeClient.Query(&q)
	curseforgeClient.Request()
	err = curseforgeClient.GetContent(&response)

	return
}
