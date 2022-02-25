package utils

import (
	"fmt"
	"github.com/lyssar/msdcli/config"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func ApiCall(requestUrl string) []byte {
	apiClient := http.Client{
		Timeout: time.Second * 10, // Timeout after 10 seconds
	}

	req, err := http.NewRequest(http.MethodGet, requestUrl, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := apiClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				log.Fatal(err)
			}
		}(res.Body)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	return body
}

func FabricApiCall(uri string, version string) []byte {
	baseUrl := config.FabricApiBaseUrl
	if version == "v2" {
		baseUrl = config.FabricApiV2BaseUrl
	}
	apiUrl := fmt.Sprintf("%s/%s", baseUrl, uri)
	return ApiCall(apiUrl)
}

func DownloadInstaller(url string) {
	fmt.Println("Downloading installer")

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Create the file
	fmt.Println("Saving file", config.InstallerFile)
	out, err := os.Create(GetCwd() + "/" + config.InstallerFile)
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
