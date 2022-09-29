package schemas

import "time"

type FileRelationType int

const (
	FileRelationTypeEmbeddedLibrary    FileRelationType = 1
	FileRelationTypeOptionalDependency FileRelationType = 2
	FileRelationTypeRequiredDependency FileRelationType = 3
	FileRelationTypeTool               FileRelationType = 4
	FileRelationTypeIncompatible       FileRelationType = 5
	FileRelationTypeInclude            FileRelationType = 6
)

type FileReleaseType int

const (
	FileReleaseTypeRelease FileReleaseType = 1
	FileReleaseTypeBeta    FileReleaseType = 2
	FileReleaseTypeAlpha   FileReleaseType = 3
)

type FileStatus int

const (
	FileStatusProcessing         FileStatus = 1
	FileStatusChangesRequired    FileStatus = 2
	FileStatusUnderReview        FileStatus = 3
	FileStatusApproved           FileStatus = 4
	FileStatusRejected           FileStatus = 5
	FileStatusMalwareDetected    FileStatus = 6
	FileStatusDeleted            FileStatus = 7
	FileStatusArchived           FileStatus = 8
	FileStatusTesting            FileStatus = 9
	FileStatusReleased           FileStatus = 10
	FileStatusReadyForReview     FileStatus = 11
	FileStatusDeprecated         FileStatus = 12
	FileStatusBaking             FileStatus = 13
	FileStatusAwaitingPublishing FileStatus = 14
	FileStatusFailedPublishing   FileStatus = 15
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
