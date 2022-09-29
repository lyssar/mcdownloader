// ----------------
// Files API @see https://docs.curseforge.com/?go#curseforge-core-api-files
// ----------------

package api

import "github.com/lyssar/msdcli/curseforge/schemas"

// GetModFile @see https://docs.curseforge.com/?go#get-mod-file
// TODO Implement
func GetModFile(modID int, fileID int) (response schemas.GetModFileResponse, err error) {
	return
}

// GetModFiles @see https://docs.curseforge.com/?go#get-mod-files
// TODO Implement
func GetModFiles(modID int, gameVersion string, modLoaderType schemas.ModLoaderType, gameVersionTypeID int, index int, pageSize int) (response schemas.GetModFilesResponse, err error) {
	return
}

// GetFiles @see https://docs.curseforge.com/?go#get-files
// TODO Implement
func GetFiles(data schemas.GetModFilesRequestBody) (response schemas.GetFilesResponse, err error) {
	return
}

// GetModFileChangelog @see https://docs.curseforge.com/?go#get-mod-file-changelog
// TODO Implement
func GetModFileChangelog(modID int, fileID int) (responses schemas.StringResponse, err error) {
	return
}

// GetModFileDownloadURL @see https://docs.curseforge.com/?go#get-mod-file-download-url
// TODO Implement
func GetModFileDownloadURL(modID int, fileID int) (responses schemas.StringResponse, err error) {
	return
}
