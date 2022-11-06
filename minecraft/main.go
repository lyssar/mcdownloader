package minecraft

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mcuadros/go-version"
	"github.com/spf13/cobra"
	"io"
	"net/http"
	"time"
)

// Filter defines filter attributes for Versions.Versions.
type Filter struct {
	Minecraft string
}

type MetaApi struct {
	ApiJSONUri string
	Versions   Versions
}

type Versions struct {
	Latest   LatestVersion `json:"latest"`
	Versions []Version     `json:"versions"`
}

type LatestVersion struct {
	Release  string `json:"release"`
	Snapshot string `json:"snapshot"`
}

type Version *struct {
	ID          string    `json:"id"`
	Type        string    `json:"type"`
	URL         string    `json:"url"`
	Time        time.Time `json:"time"`
	ReleaseTime time.Time `json:"releaseTime"`
}

func NewMinecraftMetaApi(apiJSONUri string) MetaApi {
	return MetaApi{
		ApiJSONUri: apiJSONUri,
	}
}

func (metaApi *MetaApi) LoadJson() {
	//get the data from the url
	res, err := http.Get(metaApi.ApiJSONUri)
	//error handling
	defer res.Body.Close()
	cobra.CheckErr(err)

	var versionsList Versions
	jsonBytes, _ := io.ReadAll(res.Body)
	err = json.Unmarshal([]byte(jsonBytes), &versionsList)
	cobra.CheckErr(err)
	metaApi.Versions = versionsList
}

func (metaApi *MetaApi) Filter(filter *Filter) []Version {
	var result []Version

	mc := version.NewConstrainGroupFromString(filter.Minecraft)

	for _, row := range metaApi.Versions.Versions {
		if filter.Minecraft != "" {
			if !mc.Match(row.ID) {
				continue
			}
		}

		result = append(result, row)
	}

	return result
}

func (metaApi *MetaApi) FindMinecraftVersion(selectedVersion string) (Version, error) {
	metaApi.LoadJson()

	var selectedMinecraftVersion Version
	var err error

	if selectedVersion == "" {
		selectedMinecraftVersion, err = metaApi.RenderSelect()
		cobra.CheckErr(err)
	} else {
		filter := Filter{Minecraft: selectedVersion}
		foundVersions := metaApi.Filter(&filter)
		if len(foundVersions) == 0 {
			cobra.CheckErr(
				errors.New(fmt.Sprintf("No minecraft version found for %s", selectedVersion)),
			)
		} else {
			selectedMinecraftVersion = foundVersions[0]
		}
	}
	return selectedMinecraftVersion, err
}
