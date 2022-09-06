package schemas

import "time"

type FileRelationType int

const (
	FILE_RELATION_TYPE_EMBEDDED_LIBRARY    FileRelationType = 1
	FILE_RELATION_TYPE_OPTIONAL_DEPENDENCY FileRelationType = 2
	FILE_RELATION_TYPE_REQUIRED_DEPENDENCY FileRelationType = 3
	FILE_RELATION_TYPE_TOOL                FileRelationType = 4
	FILE_RELATION_TYPE_INCOMPATIBLE        FileRelationType = 5
	FILE_RELATION_TYPE_INCLUDE             FileRelationType = 6
)

type FileReleaseType int

const (
	FILE_RELEASE_TYPE_RELEASE FileReleaseType = 1
	FILE_RELEASE_TYPE_BETA    FileReleaseType = 2
	FILE_RELEASE_TYPE_ALPHA   FileReleaseType = 3
)

type FileStatus int

const (
	FILE_STATUS_PROCESSING          FileStatus = 1
	FILE_STATUS_CHANGES_REQUIRED    FileStatus = 2
	FILE_STATUS_UNDER_REVIEW        FileStatus = 3
	FILE_STATUS_APPROVED            FileStatus = 4
	FILE_STATUS_REJECTED            FileStatus = 5
	FILE_STATUS_MALWARE_DETECTED    FileStatus = 6
	FILE_STATUS_DELETED             FileStatus = 7
	FILE_STATUS_ARCHIVED            FileStatus = 8
	FILE_STATUS_TESTING             FileStatus = 9
	FILE_STATUS_RELEASED            FileStatus = 10
	FILE_STATUS_READY_FOR_REVIEW    FileStatus = 11
	FILE_STATUS_DEPRECATED          FileStatus = 12
	FILE_STATUS_BAKING              FileStatus = 13
	FILE_STATUS_AWAITING_PUBLISHING FileStatus = 14
	FILE_STATUS_FAILED_PUBLISHING   FileStatus = 15
)

type SortableGameVersion struct {
	GameVersionName        string    `json:"gameVersionName"`
	GameVersionPadded      string    `json:"gameVersionPadded"`
	GameVersion            string    `json:"gameVersion"`
	GameVersionReleaseDate time.Time `json:"gameVersionReleaseDate"`
	GameVersionTypeID      int       `json:"gameVersionTypeId"`
}

type FileDependency struct {
	ModId        string           `json:"modId"`
	RelationType FileRelationType `json:"relationType"`
}

type FileHash struct {
	Value string   `json:"value"`
	Algo  HashAlgo `json:"algo"`
}

type FileModule struct {
	Name        string `json:"name"`
	Fingerprint int    `json:"fingerprint"`
}

type File struct {
	ID                   int                   `json:"id"`
	GameID               int                   `json:"gameId"`
	ModID                int                   `json:"modId"`
	IsAvailable          bool                  `json:"isAvailable"`
	DisplayName          string                `json:"displayName"`
	FileName             string                `json:"fileName"`
	ReleaseType          FileReleaseType       `json:"releaseType"`
	FileStatus           FileStatus            `json:"fileStatus"`
	Hashes               []FileHash            `json:"hashes"`
	FileDate             time.Time             `json:"fileDate"`
	FileLength           int                   `json:"fileLength"`
	DownloadCount        int                   `json:"downloadCount"`
	DownloadURL          string                `json:"downloadUrl"`
	GameVersions         []string              `json:"gameVersions"`
	SortableGameVersions []SortableGameVersion `json:"sortableGameVersions"`
	Dependencies         []FileDependency      `json:"dependencies"`
	ExposeAsAlternative  bool                  `json:"exposeAsAlternative"`
	ParentProjectFileID  int                   `json:"parentProjectFileId"`
	AlternateFileID      int                   `json:"alternateFileId"`
	IsServerPack         bool                  `json:"isServerPack"`
	ServerPackFileID     int                   `json:"serverPackFileId"`
	FileFingerprint      int                   `json:"fileFingerprint"`
	Modules              []FileModule          `json:"modules"`
}

type FileIndex struct {
	GameVersion       string          `json:"gameVersion"`
	FileID            int             `json:"fileId"`
	Filename          string          `json:"filename"`
	ReleaseType       FileReleaseType `json:"releaseType"`
	GameVersionTypeID int             `json:"gameVersionTypeId"`
	ModLoader         ModLoaderType   `json:"modLoader"`
}

type FingerprintFuzzyMatch struct {
	ID           int    `json:"id"`
	File         File   `json:"file"`
	LatestFiles  []File `json:"latestFiles"`
	Fingerprints []int  `json:"fingerprints"`
}

type FingerprintFuzzyMatchResult struct {
	FuzzyMatches []FingerprintFuzzyMatch `json:"fuzzyMatches"`
}

type FingerprintMatch struct {
	ID          int    `json:"id"`
	File        File   `json:"file"`
	LatestFiles []File `json:"latestFiles"`
}

type FingerprintsMatchesResult struct {
	IsCacheBuilt             bool               `json:"isCacheBuilt"`
	ExactMatches             []FingerprintMatch `json:"exactMatches"`
	ExactFingerprints        []int              `json:"exactFingerprints"`
	PartialMatches           []FingerprintMatch `json:"partialMatches"`
	PartialMatchFingerprints map[string]string  `json:"partialMatchFingerprints"`
	InstalledFingerprints    []int              `json:"installedFingerprints"`
	UnmatchedFingerprints    []int              `json:"unmatchedFingerprints"`
}

type FolderFingerprint struct {
	FolderName   string `json:"foldername"`
	Fingerprints []int  `json:"fingerprints"`
}

// RESPONSES

type GetFilesResponse struct {
	Data []File `json:"data"`
}

type GetFingerprintMatchesResponse struct {
	Data []FingerprintsMatchesResult `json:"data"`
}

type GetFingerprintsFuzzyMatchesResponse struct {
	Data []FingerprintFuzzyMatchResult `json:"data"`
}

// REQUESTS

type GetFingerprintMatchesRequestBody struct {
	Fingerprints []int `json:"fingerprints"`
}

type GetFuzzyMatchesRequestBody struct {
	GameID       int                 `json:"gameId"`
	Fingerprints []FolderFingerprint `json:"fingerprints"`
}

type GetModFilesRequestBody struct {
	FileIds []int `json:"fileIds"`
}
