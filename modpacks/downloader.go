package modpacks

import (
	"archive/zip"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/lyssar/msdcli/config"
	"github.com/manifoldco/promptui"
)

func Download() {
	var packageId int = *config.PackageId
	var serverPackageFileID int = *config.ServerPackageFileID
	var modpackDetails ModpackDetails

	if packageId <= 0 {
		validate := func(input string) error {
			_, err := strconv.ParseFloat(input, 64)
			if err != nil {
				return errors.New("Invalid project ID")
			}
			return nil
		}

		for ok := true; ok; ok = !(modpackDetails.ID > 0) {
			prompt := promptui.Prompt{
				Label:    "Enter the modpack projekt ID from curseforge",
				Validate: validate,
			}

			result, err := prompt.Run()

			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}
			packageId, err = strconv.Atoi(result)

			modpackDetails = fetchDetailJson(packageId)
			if !(modpackDetails.ID > 0) {
				fmt.Printf("Project with ID %d doesn't exist on curseforge.\n", packageId)
			}
		}
	} else {
		modpackDetails = fetchDetailJson(packageId)
		if !(modpackDetails.ID > 0) {
			fmt.Printf("Project with ID %d doesn't exist on curseforge.\n", packageId)
			os.Exit(1)
		}
	}

	var modpackVersion ModpackFile

	if serverPackageFileID <= 0 {
		modpackFiles := fetchVersionsOfModpack(modpackDetails.ID)
		modpackVersion = chooseModpackVersion(modpackFiles)
	} else {
		modpackVersion = ModpackFile{ServerPackFileID: serverPackageFileID, ID: packageId}
	}

	var modpackServerFile string
	if modpackVersion.ServerPackFileID > 0 {
		modpackServerFile = fetchServerFileUrl(modpackVersion.ID, modpackVersion.ServerPackFileID)
	} else {
		modpackServerFile = modpackVersion.DownloadURL
	}

	downloadZipFile(modpackServerFile)

	unzipModpack("server.zip", true)
	e := os.Remove("server.zip")
	if e != nil {
		log.Fatal(e)
	}
}

func chooseModpackVersion(modpackFiles ModpackFiles) ModpackFile {
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   promptui.IconSelect + " {{ .DisplayName | cyan }}",
		Inactive: "  {{ .DisplayName | white }}",
		Selected: promptui.IconGood + " Modpack version: {{ .DisplayName | black }}",
	}

	prompt := promptui.Select{
		Label:     "Select modpack version:",
		Items:     modpackFiles,
		Templates: templates,
	}

	i, _, err := prompt.Run()

	if err != nil {
		log.Fatal(err)
	}

	return modpackFiles[i]
}

func fetchServerFileUrl(modpackID int, serverPackFileID int) string {
	requestUrl := fmt.Sprintf("%s/addon/%d/file/%d/download-url", config.CursforgeApiUrl, modpackID, serverPackFileID)
	body := apiCall(requestUrl)

	return string(body)
}

func fetchDetailJson(packageId int) ModpackDetails {
	requestUrl := fmt.Sprintf("%s/addon/%d", config.CursforgeApiUrl, packageId)
	body := apiCall(requestUrl)

	modpackDetails := ModpackDetails{}
	jsonErr := json.Unmarshal(body, &modpackDetails)
	if jsonErr != nil {
		return modpackDetails
	}

	return modpackDetails
}

func fetchVersionsOfModpack(packageId int) ModpackFiles {
	requestUrl := fmt.Sprintf("%s/addon/%d/files", config.CursforgeApiUrl, packageId)
	body := apiCall(requestUrl)

	modpackFiles := ModpackFiles{}
	jsonErr := json.Unmarshal(body, &modpackFiles)
	if jsonErr != nil {
		return modpackFiles
	}

	var filteredModpackFiles ModpackFiles
	for _, v := range modpackFiles {
		if v.ReleaseType == 1 {
			filteredModpackFiles = append(filteredModpackFiles, v)
		}
	}

	return filteredModpackFiles
}

func apiCall(requestUrl string) []byte {
	curseforgeApiClient := http.Client{
		Timeout: time.Second * 10, // Timeout after 10 seconds
	}

	req, err := http.NewRequest(http.MethodGet, requestUrl, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := curseforgeApiClient.Do(req)
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

func unzipModpack(file string, cutRoot bool) {
	dst, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	archive, err := zip.OpenReader(file)
	if err != nil {
		panic(err)
	}
	defer archive.Close()

	rootElements := 0
	lastFolder := ""
	if cutRoot == true {
		for _, tmpf := range archive.File {
			parts := strings.Split(tmpf.Name, "/")
			if len(parts) == 1 {
				rootElements = rootElements + 1
			} else if len(parts) > 1 {
				if lastFolder != parts[0] {
					rootElements = rootElements + 1
					lastFolder = parts[0]
				}
			}
		}
	}

	for _, f := range archive.File {
		fileName := f.Name
		parts := strings.Split(fileName, "/")
		if cutRoot == true && rootElements <= 1 && len(parts) > 1 {
			parts = parts[1:]
			fileName = strings.Join(parts, "/")
		}

		filePath := filepath.Join(dst, fileName)
		fmt.Println("unzipping file ", filePath)

		if filePath == filepath.Clean(dst) {
			continue
		}

		if !strings.HasPrefix(filePath, filepath.Clean(dst)+string(os.PathSeparator)) {
			fmt.Println("invalid file path")
			return
		}
		if f.FileInfo().IsDir() {
			fmt.Println("creating directory...")
			os.MkdirAll(filePath, os.ModePerm)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			panic(err)
		}

		dstFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			panic(err)
		}

		fileInArchive, err := f.Open()
		if err != nil {
			panic(err)
		}

		if _, err := io.Copy(dstFile, fileInArchive); err != nil {
			panic(err)
		}

		dstFile.Close()
		fileInArchive.Close()
	}
}

func downloadZipFile(file string) {
	resp, err := http.Get(file)
	if err != nil {
		fmt.Printf("err: %s", err)
	}

	defer resp.Body.Close()
	fmt.Println("status", resp.Status)
	if resp.StatusCode != 200 {
		return
	}

	// Create the file
	out, err := os.Create("server.zip")
	if err != nil {
		fmt.Printf("err: %s", err)
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	fmt.Printf("err: %s", err)
}
