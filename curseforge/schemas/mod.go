package schemas

import "time"

type ModStatus int

const (
	ModStatusNew             ModStatus = 1
	ModStatusChangesRequired ModStatus = 2
	ModStatusUnderSoftReview ModStatus = 3
	ModStatusApproved        ModStatus = 4
	ModStatusRejected        ModStatus = 5
	ModStatusChangesMade     ModStatus = 6
	ModStatusInactive        ModStatus = 7
	ModStatusAbandoned       ModStatus = 8
	ModStatusDeleted         ModStatus = 9
	ModStatusUnderReview     ModStatus = 10
)

type ModsSearchSortField int

const (
	ModSearchSortFieldFeatured       ModsSearchSortField = 1
	ModSearchSortFieldPopularity     ModsSearchSortField = 2
	ModSearchSortFieldLastUpdated    ModsSearchSortField = 3
	ModSearchSortFieldName           ModsSearchSortField = 4
	ModSearchSortFieldAuthor         ModsSearchSortField = 5
	ModSearchSortFieldTotalDownloads ModsSearchSortField = 6
	ModSearchSortFieldCategory       ModsSearchSortField = 7
	ModSearchSortFieldGameVersion    ModsSearchSortField = 8
)

type SortOrder string

const (
	SortOrderAsc  SortOrder = "asc"
	SortOrderDesc SortOrder = "desc"
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
