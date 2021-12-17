package server

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
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

func Forge() {
	fmt.Println("Start Installing MC Server")

	minecraftVersion := selectMinecraftVersion()

	fmt.Printf("MC Version: %s\n", minecraftVersion.Version)

	forgeVersion := selectForgeVersion(minecraftVersion)

	fmt.Printf("Forge Version: %s\n", forgeVersion.Version)

	downloadForge(forgeVersion)
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

func downloadForge(forgeVersion ForgeVersion) error {
	fmt.Println("Downloading forge version: ", forgeVersion.Version)

	// Get the data
	resp, err := http.Get(forgeVersion.Installer)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(getCwd() + "/forger_installer.jar")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func getCwd() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return dir
}
