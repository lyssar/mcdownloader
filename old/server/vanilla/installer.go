package vanilla

import (
	"encoding/json"
	"fmt"
	"github.com/lyssar/msdcli/config"
	"github.com/lyssar/msdcli/utils"
	"io"
	"log"
	"net/http"
	"os"
)

func DownloadInstaller() {
	var minecraftVersion utils.MinecraftVersion
	if *config.McVersion == "" {
		minecraftVersion = utils.SelectMinecraftVersion()
	} else {
		minecraftVersion = utils.GetMinecraftVersionInfo(*config.McVersion)
	}
	// downloads
	fmt.Printf("%+v\n", minecraftVersion)
	downloadServerJar(minecraftVersion)
}

func InstalServer() {
	utils.CreateEula()
	fmt.Println("Vanilla server installed. Please create and configure your server.properties file before starting")
}

func downloadServerJar(mcVersion utils.MinecraftVersion) {
	fmt.Println("Downloading minecraft vanilla version: ", mcVersion.ID)
	serverDownloadJar := fetchDownloadURL(mcVersion.URL)

	// Get the data
	resp, err := http.Get(serverDownloadJar)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Create the file
	fmt.Println("Loading", config.ServerFile)
	out, err := os.Create(utils.GetCwd() + "/" + config.ServerFile)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
}

func fetchDownloadURL(detailJson string) string {
	var downloadUrl string
	var minecraftVersionDetails utils.MinecraftVersionDetails

	requestUrl := fmt.Sprintf(detailJson)
	body := utils.ApiCall(requestUrl)

	jsonErr := json.Unmarshal(body, &minecraftVersionDetails)
	if jsonErr != nil {
		return downloadUrl
	}

	return minecraftVersionDetails.Downloads.Server.URL
}
