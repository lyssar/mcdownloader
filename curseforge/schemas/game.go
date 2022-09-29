package schemas

import "time"

type GameVersionStatus int

const (
	GameVersionStatusApproved GameVersionStatus = 1
	GameVersionStatusDeleted  GameVersionStatus = 2
	GameVersionStatusNew      GameVersionStatus = 3
)

type GameVersionTypeStatus int

const (
	GameVersionTypeStatusNormal  GameVersionTypeStatus = 1
	GameVersionTypeStatusDeleted GameVersionTypeStatus = 2
)

type GameAssets struct {
	IconURL  string `json:"iconUrl"`
	TileURL  string `json:"tileUrl"`
	CoverURL string `json:"coverUrl"`
}

type Game struct {
	ID           int           `json:"id"`
	Name         string        `json:"name"`
	Slug         string        `json:"slug"`
	DateModified time.Time     `json:"dateModified"`
	Assets       GameAssets    `json:"assets"`
	Status       CoreStatus    `json:"status"`
	APIStatus    CoreApiStatus `json:"apiStatus"`
}

type GameVersionsByType struct {
	Type     int      `json:"type"`
	Versions []string `json:"versions"`
}

type GameVersionType struct {
	ID     int    `json:"id"`
	GameID int    `json:"gameId"`
	Name   string `json:"name"`
	Slug   string `json:"slug"`
}

// RESPONSES

type GetGameResponse struct {
	Data Game `json:"data"`
}

type GetGamesResponse struct {
	Data       []Game       `json:"data"`
	Pagination []Pagination `json:"pagination"`
}

type GetVersionTypesResponse struct {
	Data []GameVersionType `json:"data"`
}

type GetVersionsResponse struct {
	Data []GameVersionsByType `json:"data"`
}
