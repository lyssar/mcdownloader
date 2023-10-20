package minecraft

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"github.com/lyssar/msdcli/errors"
	"github.com/lyssar/msdcli/logger"
	"github.com/lyssar/msdcli/utils"
	"github.com/mcuadros/go-version"
	"io"
	"net/http"
	"os"
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

//go:embed resources/eula.txt
var defaultEula []byte

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
	logger.Info("Loading minecraft meta data")
	//get the data from the url
	logger.Debug(metaApi.ApiJSONUri)
	res, err := http.Get(metaApi.ApiJSONUri)
	//error handling
	defer res.Body.Close()
	errors.CheckStandardErr(err)

	var versionsList Versions
	jsonBytes, _ := io.ReadAll(res.Body)
	err = json.Unmarshal([]byte(jsonBytes), &versionsList)
	errors.CheckStandardErr(err)
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

func (metaApi *MetaApi) FindMinecraftVersion(selectedVersion string) (*Version, *errors.ApplicationError) {
	logger.Info("Choosing minecraft version")
	metaApi.LoadJson()
	if selectedVersion == "" {
		selectedMinecraftVersion, err := metaApi.RenderSelect(false)
		if err != nil {
			return nil, errors.NewError(err.Error())
		}
		return &selectedMinecraftVersion, nil
	} else {
		filter := Filter{Minecraft: selectedVersion}
		foundVersions := metaApi.Filter(&filter)
		if len(foundVersions) == 0 {
			return nil, errors.NewError(fmt.Sprintf("No minecraft version found for %s", selectedVersion))
		} else {
			selectedMinecraftVersion := foundVersions[0]
			return &selectedMinecraftVersion, nil
		}
	}
}

func (versionMeta Version) DownloadRelease() McRelease {
	mcRelease := McRelease{}
	req, _ := http.NewRequest("GET", versionMeta.URL, nil)
	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()

	responseData, err := io.ReadAll(resp.Body)
	errors.CheckStandardErr(err)

	_ = json.Unmarshal([]byte(responseData), &mcRelease)

	return mcRelease
}

func (r McRelease) DownloadServer(workingDir string) bool {
	if _, err := os.Stat(workingDir); os.IsNotExist(err) {
		errors.CheckStandardErr(err)
	}
	utils.DownloadFile(r.Downloads.Server.URL, "server.jar", workingDir)
	return true
}

func (r McRelease) InstallServer(workingDir string) {
	if _, err := os.Stat(workingDir); os.IsNotExist(err) {
		errors.CheckStandardErr(err)
	}

	srvPath := fmt.Sprintf("%s/server.jar", workingDir)
	_, err := os.Stat(srvPath)

	if os.IsNotExist(err) {
		errors.CheckStandardErr(err)
	}

	eulaFile := fmt.Sprintf("%s/eula.txt", workingDir)
	_, err = os.Stat(eulaFile)

	if os.IsNotExist(err) {
		// Create the file
		f, err := os.Create(eulaFile)
		defer f.Close()
		errors.CheckStandardErr(err)
		_, err = f.WriteString(string(defaultEula))
		errors.CheckStandardErr(err)
		logger.Success("Eula accepted")
	}
}
