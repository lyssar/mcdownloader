package schemas

import "time"

type Category struct {
	ID               int       `json:"id"`
	GameID           int       `json:"gameId"`
	Name             string    `json:"name"`
	Slug             string    `json:"slug"`
	URL              string    `json:"url"`
	IconURL          string    `json:"iconUrl"`
	DateModified     time.Time `json:"dateModified"`
	IsClass          *bool     `json:"isClass"`
	ClassID          *int      `json:"classId"`
	ParentCategoryID *int      `json:"parentCategoryId"`
	DisplayIndex     *int      `json:"displayIndex"`
}

// RESPONSES

type GetCategoriesResponse struct {
	Data []Category `json:"data"`
}
