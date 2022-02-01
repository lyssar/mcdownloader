package utils

import "time"

type MinecraftVersion struct {
	ID              string    `json:"id"`
	Type            string    `json:"type"`
	URL             string    `json:"url"`
	Time            time.Time `json:"time"`
	ReleaseTime     time.Time `json:"releaseTime"`
	Sha1            string    `json:"sha1"`
	ComplianceLevel int       `json:"complianceLevel"`
}

type ModpackFile struct {
	ID              int           `json:"id"`
	DisplayName     string        `json:"displayName"`
	FileName        string        `json:"fileName"`
	FileDate        time.Time     `json:"fileDate"`
	FileLength      int           `json:"fileLength"`
	ReleaseType     int           `json:"releaseType"`
	FileStatus      int           `json:"fileStatus"`
	DownloadURL     string        `json:"downloadUrl"`
	IsAlternate     bool          `json:"isAlternate"`
	AlternateFileID int           `json:"alternateFileId"`
	Dependencies    []interface{} `json:"dependencies"`
	IsAvailable     bool          `json:"isAvailable"`
	Modules         []struct {
		Foldername  string `json:"foldername"`
		Fingerprint int64  `json:"fingerprint"`
	} `json:"modules"`
	PackageFingerprint      int         `json:"packageFingerprint"`
	GameVersion             []string    `json:"gameVersion"`
	InstallMetadata         interface{} `json:"installMetadata"`
	ServerPackFileID        int         `json:"serverPackFileId"`
	HasInstallScript        bool        `json:"hasInstallScript"`
	GameVersionDateReleased time.Time   `json:"gameVersionDateReleased"`
	GameVersionFlavor       interface{} `json:"gameVersionFlavor"`
}

type ModpackFiles []ModpackFile

type ModpackDetails struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Authors []struct {
		Name              string      `json:"name"`
		URL               string      `json:"url"`
		ProjectID         int         `json:"projectId"`
		ID                int         `json:"id"`
		ProjectTitleID    interface{} `json:"projectTitleId"`
		ProjectTitleTitle interface{} `json:"projectTitleTitle"`
		UserID            int         `json:"userId"`
		TwitchID          int         `json:"twitchId"`
	} `json:"authors"`
	Attachments []struct {
		ID           int    `json:"id"`
		ProjectID    int    `json:"projectId"`
		Description  string `json:"description"`
		IsDefault    bool   `json:"isDefault"`
		ThumbnailURL string `json:"thumbnailUrl"`
		Title        string `json:"title"`
		URL          string `json:"url"`
		Status       int    `json:"status"`
	} `json:"attachments"`
	WebsiteURL    string  `json:"websiteUrl"`
	GameID        int     `json:"gameId"`
	Summary       string  `json:"summary"`
	DefaultFileID int     `json:"defaultFileId"`
	DownloadCount float64 `json:"downloadCount"`
	LatestFiles   []struct {
		ID              int           `json:"id"`
		DisplayName     string        `json:"displayName"`
		FileName        string        `json:"fileName"`
		FileDate        time.Time     `json:"fileDate"`
		FileLength      int           `json:"fileLength"`
		ReleaseType     int           `json:"releaseType"`
		FileStatus      int           `json:"fileStatus"`
		DownloadURL     string        `json:"downloadUrl"`
		IsAlternate     bool          `json:"isAlternate"`
		AlternateFileID int           `json:"alternateFileId"`
		Dependencies    []interface{} `json:"dependencies"`
		IsAvailable     bool          `json:"isAvailable"`
		Modules         []struct {
			Foldername  string `json:"foldername"`
			Fingerprint int64  `json:"fingerprint"`
			Type        int    `json:"type"`
		} `json:"modules"`
		PackageFingerprint  int      `json:"packageFingerprint"`
		GameVersion         []string `json:"gameVersion"`
		SortableGameVersion []struct {
			GameVersionPadded      string    `json:"gameVersionPadded"`
			GameVersion            string    `json:"gameVersion"`
			GameVersionReleaseDate time.Time `json:"gameVersionReleaseDate"`
			GameVersionName        string    `json:"gameVersionName"`
		} `json:"sortableGameVersion"`
		InstallMetadata            interface{}   `json:"installMetadata"`
		Changelog                  interface{}   `json:"changelog"`
		HasInstallScript           bool          `json:"hasInstallScript"`
		IsCompatibleWithClient     bool          `json:"isCompatibleWithClient"`
		CategorySectionPackageType int           `json:"categorySectionPackageType"`
		RestrictProjectFileAccess  int           `json:"restrictProjectFileAccess"`
		ProjectStatus              int           `json:"projectStatus"`
		RenderCacheID              int           `json:"renderCacheId"`
		FileLegacyMappingID        interface{}   `json:"fileLegacyMappingId"`
		ProjectID                  int           `json:"projectId"`
		ParentProjectFileID        interface{}   `json:"parentProjectFileId"`
		ParentFileLegacyMappingID  interface{}   `json:"parentFileLegacyMappingId"`
		FileTypeID                 interface{}   `json:"fileTypeId"`
		ExposeAsAlternative        interface{}   `json:"exposeAsAlternative"`
		PackageFingerprintID       int           `json:"packageFingerprintId"`
		GameVersionDateReleased    time.Time     `json:"gameVersionDateReleased"`
		GameVersionMappingID       int           `json:"gameVersionMappingId"`
		GameVersionID              int           `json:"gameVersionId"`
		GameID                     int           `json:"gameId"`
		IsServerPack               bool          `json:"isServerPack"`
		ServerPackFileID           int           `json:"serverPackFileId"`
		GameVersionFlavor          interface{}   `json:"gameVersionFlavor"`
		Hashes                     []interface{} `json:"hashes"`
		DownloadCount              int           `json:"downloadCount"`
	} `json:"latestFiles"`
	Categories []struct {
		CategoryID   int       `json:"categoryId"`
		Name         string    `json:"name"`
		URL          string    `json:"url"`
		AvatarURL    string    `json:"avatarUrl"`
		ParentID     int       `json:"parentId"`
		RootID       int       `json:"rootId"`
		ProjectID    int       `json:"projectId"`
		AvatarID     int       `json:"avatarId"`
		GameID       int       `json:"gameId"`
		Slug         string    `json:"slug"`
		DateModified time.Time `json:"dateModified"`
	} `json:"categories"`
	Status            int `json:"status"`
	PrimaryCategoryID int `json:"primaryCategoryId"`
	CategorySection   struct {
		ID                      int         `json:"id"`
		GameID                  int         `json:"gameId"`
		Name                    string      `json:"name"`
		PackageType             int         `json:"packageType"`
		Path                    string      `json:"path"`
		InitialInclusionPattern string      `json:"initialInclusionPattern"`
		ExtraIncludePattern     interface{} `json:"extraIncludePattern"`
		GameCategoryID          int         `json:"gameCategoryId"`
	} `json:"categorySection"`
	Slug                   string `json:"slug"`
	GameVersionLatestFiles []struct {
		GameVersion       string      `json:"gameVersion"`
		ProjectFileID     int         `json:"projectFileId"`
		ProjectFileName   string      `json:"projectFileName"`
		FileType          int         `json:"fileType"`
		GameVersionFlavor interface{} `json:"gameVersionFlavor"`
	} `json:"gameVersionLatestFiles"`
	IsFeatured           bool        `json:"isFeatured"`
	PopularityScore      float64     `json:"popularityScore"`
	GamePopularityRank   int         `json:"gamePopularityRank"`
	PrimaryLanguage      string      `json:"primaryLanguage"`
	GameSlug             string      `json:"gameSlug"`
	GameName             string      `json:"gameName"`
	PortalName           string      `json:"portalName"`
	DateModified         time.Time   `json:"dateModified"`
	DateCreated          time.Time   `json:"dateCreated"`
	DateReleased         time.Time   `json:"dateReleased"`
	IsAvailable          bool        `json:"isAvailable"`
	IsExperiemental      bool        `json:"isExperiemental"`
	AllowModDistribution interface{} `json:"allowModDistribution"`
}

type MinecraftVersionList struct {
	Latest struct {
		Release  string `json:"release"`
		Snapshot string `json:"snapshot"`
	} `json:"latest"`
	Versions []MinecraftVersion `json:"versions"`
}
