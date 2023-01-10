package minecraft

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gookit/color"
	"github.com/lyssar/msdcli/utils"
	"github.com/mcuadros/go-version"
	"github.com/spf13/cobra"
	"io"
	"net/http"
	"os"
	"os/exec"
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

func (v Versions) GetVersionsForType(releaseType string) []Version {
	var filteredList []Version
	//filteredList = append(filteredList, )
	for _, mcVersion := range v.Versions {
		if mcVersion.Type == releaseType {
			filteredList = append(filteredList, mcVersion)
		}
	}

	return filteredList
}

type LatestVersion struct {
	Release  string `json:"release"`
	Snapshot string `json:"snapshot"`
}

type Version struct {
	ID          string    `json:"id"`
	Type        string    `json:"type"`
	URL         string    `json:"url"`
	Time        time.Time `json:"time"`
	ReleaseTime time.Time `json:"releaseTime"`
}

type McRelease struct {
	Arguments struct {
		Game []interface{} `json:"game"`
		Jvm  []interface{} `json:"jvm"`
	} `json:"arguments"`
	AssetIndex struct {
		ID        string `json:"id"`
		Sha1      string `json:"sha1"`
		Size      int    `json:"size"`
		TotalSize int    `json:"totalSize"`
		URL       string `json:"url"`
	} `json:"assetIndex"`
	Assets          string `json:"assets"`
	ComplianceLevel int    `json:"complianceLevel"`
	Downloads       struct {
		Client struct {
			Sha1 string `json:"sha1"`
			Size int    `json:"size"`
			URL  string `json:"url"`
		} `json:"client"`
		ClientMappings struct {
			Sha1 string `json:"sha1"`
			Size int    `json:"size"`
			URL  string `json:"url"`
		} `json:"client_mappings"`
		Server struct {
			Sha1 string `json:"sha1"`
			Size int    `json:"size"`
			URL  string `json:"url"`
		} `json:"server"`
		ServerMappings struct {
			Sha1 string `json:"sha1"`
			Size int    `json:"size"`
			URL  string `json:"url"`
		} `json:"server_mappings"`
	} `json:"downloads"`
	ID          string `json:"id"`
	JavaVersion struct {
		Component    string `json:"component"`
		MajorVersion int    `json:"majorVersion"`
	} `json:"javaVersion"`
	Libraries []struct {
		Downloads struct {
			Artifact struct {
				Path string `json:"path"`
				Sha1 string `json:"sha1"`
				Size int    `json:"size"`
				URL  string `json:"url"`
			} `json:"artifact"`
		} `json:"downloads"`
		Name  string `json:"name"`
		Rules []struct {
			Action string `json:"action"`
			Os     struct {
				Name string `json:"name"`
			} `json:"os"`
		} `json:"rules,omitempty"`
	} `json:"libraries"`
	Logging struct {
		Client struct {
			Argument string `json:"argument"`
			File     struct {
				ID   string `json:"id"`
				Sha1 string `json:"sha1"`
				Size int    `json:"size"`
				URL  string `json:"url"`
			} `json:"file"`
			Type string `json:"type"`
		} `json:"client"`
	} `json:"logging"`
	MainClass              string    `json:"mainClass"`
	MinimumLauncherVersion int       `json:"minimumLauncherVersion"`
	ReleaseTime            time.Time `json:"releaseTime"`
	Time                   time.Time `json:"time"`
	Type                   string    `json:"type"`
}

func NewMinecraftMetaApi(apiJSONUri string) MetaApi {
	return MetaApi{
		ApiJSONUri: apiJSONUri,
	}
}

func (metaApi *MetaApi) LoadJson() {
	utils.PrintInfo("Loading minecraft meta data")
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
	utils.PrintInfo("Choosing minecraft version")
	metaApi.LoadJson()

	var selectedMinecraftVersion Version
	var err error

	if selectedVersion == "" {
		selectedMinecraftVersion, err = metaApi.RenderSelect(false)
		cobra.CheckErr(err)
	} else {
		filter := Filter{Minecraft: selectedVersion}
		foundVersions := metaApi.Filter(&filter)
		if len(foundVersions) == 0 {
			cobra.CheckErr(
				errors.New(color.Error.Sprintf("No minecraft version found for %s", selectedVersion)),
			)
		} else {
			selectedMinecraftVersion = foundVersions[0]
		}
	}
	return selectedMinecraftVersion, err
}

func (versionMeta Version) DownloadRelease() McRelease {
	mcRelease := McRelease{}
	req, _ := http.NewRequest("GET", versionMeta.URL, nil)
	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()

	responseData, err := io.ReadAll(resp.Body)
	cobra.CheckErr(err)

	_ = json.Unmarshal([]byte(responseData), &mcRelease)

	return mcRelease
}

func (r McRelease) DownloadServer() bool {
	workingDir, err := os.Getwd()
	cobra.CheckErr(err)
	utils.DownloadFile(r.Downloads.Server.URL, "server.jar", workingDir)
	return true
}

func (r McRelease) InstallServer() {
	pwd, err := os.Getwd()
	cobra.CheckErr(err)

	srvPath := fmt.Sprintf("%s/server.jar", pwd)
	_, err = os.Stat(srvPath)

	if os.IsNotExist(err) {
		cobra.CheckErr(err)
	}

	cmd := exec.Command("java", "-Xms2048M", "-Xmx2048M", "-jar", srvPath, "nogui")
	err = cmd.Start()
	cobra.CheckErr(err)
}
