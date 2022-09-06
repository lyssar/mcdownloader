package schemas

import "time"

type ModStatus int

const (
	MOD_STATUS_NEW               ModStatus = 1
	MOD_STATUS_CHANGES_REQUIRED  ModStatus = 2
	MOD_STATUS_UNDER_SOFT_REVIEW ModStatus = 3
	MOD_STATUS_APPROVED          ModStatus = 4
	MOD_STATUS_REJECTED          ModStatus = 5
	MOD_STATUS_CHANGES_MADE      ModStatus = 6
	MOD_STATUS_INACTIVE          ModStatus = 7
	MOD_STATUS_ABANDONED         ModStatus = 8
	MOD_STATUS_DELETED           ModStatus = 9
	MOD_STATUS_UNDER_REVIEW      ModStatus = 10
)

type ModsSearchSortField int

const (
	FILE_STATUS_FEATURED        FileStatus = 1
	FILE_STATUS_POPULARITY      FileStatus = 2
	FILE_STATUS_LAST_UPDATED    FileStatus = 3
	FILE_STATUS_NAME            FileStatus = 4
	FILE_STATUS_AUTHOR          FileStatus = 5
	FILE_STATUS_TOTAL_DOWNLOADS FileStatus = 6
	FILE_STATUS_CATEGORY        FileStatus = 7
	FILE_STATUS_GAME_VERSION    FileStatus = 8
)

type SortOrder string

const (
	SORT_ORDER_ASC  SortOrder = "asc"
	SORT_ORDER_DESC SortOrder = "asc"
)

type FeaturedModsResponse struct {
	Featured        []Mod `json:"featured"`
	Popular         []Mod `json:"popular"`
	RecentlyUpdated []Mod `json:"recentlyUpdated"`
}

type ModLinks struct {
	WebsiteURL string `json:"websiteUrl"`
	WikiURL    string `json:"wikiUrl"`
	IssuesURL  string `json:"issuesUrl"`
	SourceURL  string `json:"sourceUrl"`
}

type ModAuthor struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

type ModAsset struct {
	ID           int    `json:"id"`
	ModID        int    `json:"modId"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	ThumbnailURL string `json:"thumbnailUrl"`
	URL          string `json:"url"`
}

type Mod struct {
	ID                   int         `json:"id"`
	GameID               int         `json:"gameId"`
	Name                 string      `json:"name"`
	Slug                 string      `json:"slug"`
	Links                ModLinks    `json:"links"`
	Summary              string      `json:"summary"`
	Status               ModStatus   `json:"status"`
	DownloadCount        int         `json:"downloadCount"`
	IsFeatured           bool        `json:"isFeatured"`
	PrimaryCategoryID    int         `json:"primaryCategoryId"`
	Categories           []Category  `json:"categories"`
	ClassID              *int        `json:"classId"`
	Authors              []ModAuthor `json:"authors"`
	Logo                 ModAsset    `json:"logo"`
	Screenshots          []ModAsset  `json:"screenshots"`
	MainFileID           int         `json:"mainFileId"`
	LatestFiles          []File      `json:"latestFiles"`
	LatestFilesIndexes   []FileIndex `json:"latestFilesIndexes"`
	DateCreated          time.Time   `json:"dateCreated"`
	DateModified         time.Time   `json:"dateModified"`
	DateReleased         time.Time   `json:"dateReleased"`
	AllowModDistribution *bool       `json:"allowModDistribution"`
	GamePopularityRank   int         `json:"gamePopularityRank"`
	IsAvailable          bool        `json:"isAvailable"`
	ThumbsUpCount        int         `json:"thumbsUpCount"`
}

// RESPONSES

type GetFeaturedModsResponse struct {
	Data []FeaturedModsResponse `json:"data"`
}

type GetModFileResponse struct {
	Data File `json:"data"`
}

type GetModFilesResponse struct {
	Data       []File     `json:"data"`
	Pagination Pagination `json:"pagination"`
}

type GetModResponse struct {
	Data Mod `json:"data"`
}

type GetModsResponse struct {
	Data []Mod `json:"data"`
}

type SearchModsResponse struct {
	Data       []Mod      `json:"data"`
	Pagination Pagination `json:"pagination"`
}

type StringResponse struct {
	Data string `json:"data"`
}

// REQUESTS

type GetFeaturedModsRequestBody struct {
	GameID            int   `json:"gameId"`
	ExcludedModIds    []int `json:"excludedModIds"`
	GameVersionTypeID int   `json:"gameVersionTypeId"`
}

type GetModsByIdsListRequestBody struct {
	ModIds []int `json:"modIds"`
}
