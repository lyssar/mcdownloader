package forge

import (
	"fmt"
	"github.com/lyssar/mcdownloader/config"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/antchfx/htmlquery"
	"github.com/manifoldco/promptui"
)

type ForgeVersion struct {
	Version   string
	Installer string
}

type MinecraftVersion struct {
	Version string
	Page    string
}

func DownloadInstaller() (MinecraftVersion, ForgeVersion) {
	minecraftVersion := selectMinecraftVersion()

	fmt.Printf("MC Version: %s\n", minecraftVersion.Version)

	forgeVersion := selectForgeVersion(minecraftVersion)

	fmt.Printf("Forge Version: %s\n", forgeVersion.Version)

	downloadForge(forgeVersion)

	eulaFile := []byte("forge-" + minecraftVersion.Version + "-" + forgeVersion.Version)
	err := os.WriteFile("version.txt", eulaFile, 0644)

	if err != nil {
		log.Fatal(err)
	}

	return minecraftVersion, forgeVersion
}

func InstalServer(minecraftVerion MinecraftVersion, forgeVersion ForgeVersion) {
	eulaFile := []byte("eula=true")
	err := os.WriteFile("eula.txt", eulaFile, 0644)

	if err != nil {
		log.Fatal(err)
	}

	out, cmdErr := exec.Command("/usr/bin/java", "-jar", getCwd()+"/"+config.InstallerFile, "--installServer").Output()

	if cmdErr != nil {
		log.Fatal(cmdErr)
	}
	fmt.Println(string(out))

	renamErr := os.Rename("forge-"+minecraftVerion.Version+"-"+forgeVersion.Version+".jar", config.ServerFile)
	if renamErr != nil {
		log.Fatal(renamErr)
	}

	err = os.Remove(config.InstallerFile)
	if err != nil {
		return
	}

	err = os.Remove(config.InstallerFile + ".log")
	if err != nil {
		return
	}

	fmt.Println("Forge installed. Please configure your server.properties file before starting")
}

func selectMinecraftVersion() MinecraftVersion {
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   promptui.IconSelect + " {{ .Version | cyan }}",
		Inactive: "  {{ .Version | white }}",
		Selected: promptui.IconGood + " Mincraft Version: {{ .Version | black }}",
	}

	versionList := getMinecraftVersionList()

	prompt := promptui.Select{
		Label:     "Select Minecraft version:",
		Items:     versionList,
		Templates: templates,
	}

	i, _, err := prompt.Run()

	if err != nil {
		log.Fatal(err)
	}

	return versionList[i]
}

func selectForgeVersion(minecraftVersion MinecraftVersion) ForgeVersion {
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   promptui.IconSelect + " {{ .Version | cyan }}",
		Inactive: "  {{ .Version | white }}",
		Selected: promptui.IconGood + " Mincraft Version: {{ .Version | black }}",
	}

	forgeVersionList := getForgeVersionForMinecraftVersion(minecraftVersion)

	prompt := promptui.Select{
		Label:     "Select forge version:",
		Items:     forgeVersionList,
		Templates: templates,
	}

	i, _, err := prompt.Run()

	if err != nil {
		log.Fatal(err)
	}

	return forgeVersionList[i]
}

func getMinecraftVersionList() []MinecraftVersion {

	versionListItems := []MinecraftVersion{}

	doc, err := htmlquery.LoadURL("https://files.minecraftforge.net/net/minecraftforge/forge/")
	if err != nil {
		log.Fatal(err)
	}

	activeVersion := htmlquery.FindOne(doc, "//li[contains(@class, 'li-version-list')]/ul/li[contains(@class, 'elem-active')]")
	minecraftVersion := MinecraftVersion{Version: htmlquery.InnerText(activeVersion), Page: "index_" + htmlquery.InnerText(activeVersion) + ".html"}

	versionListItems = append(versionListItems, minecraftVersion)

	links := htmlquery.Find(doc, "//li[contains(@class, 'li-version-list')]/ul/li/a[@href]")
	for _, n := range links {
		url := htmlquery.SelectAttr(n, "href")
		versionStr := strings.ReplaceAll(strings.ReplaceAll(url, "index_", ""), ".html", "")

		minecraftVersion := MinecraftVersion{Version: versionStr, Page: url}

		versionListItems = append(versionListItems, minecraftVersion)
	}

	return versionListItems
}

func getForgeVersionForMinecraftVersion(mcVersion MinecraftVersion) []ForgeVersion {
	versionList := []ForgeVersion{}

	doc, err := htmlquery.LoadURL("https://files.minecraftforge.net/net/minecraftforge/forge/" + mcVersion.Page)
	if err != nil {
		log.Fatal(err)
	}

	links := htmlquery.Find(doc, "//table[contains(@class, 'download-list')]/tbody/tr/td[contains(@class, 'download-version')]")
	for _, n := range links {
		forgeVersion := strings.Trim(htmlquery.InnerText(n), " \n")
		installer := fmt.Sprintf("https://maven.minecraftforge.net/net/minecraftforge/forge/%[1]s-%[2]s/forge-%[1]s-%[2]s-installer.jar", mcVersion.Version, forgeVersion)

		forgetVersion := ForgeVersion{Version: forgeVersion, Installer: installer}

		versionList = append(versionList, forgetVersion)
	}

	return versionList
}

func downloadForge(forgeVersion ForgeVersion) {
	fmt.Println("Downloading forge version: ", forgeVersion.Version)

	// Get the data
	resp, err := http.Get(forgeVersion.Installer)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Create the file
	fmt.Println("Loading", config.InstallerFile)
	out, err := os.Create(getCwd() + "/" + config.InstallerFile)
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

func getCwd() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return dir
}
