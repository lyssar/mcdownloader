package forge

import (
	"fmt"
	"github.com/lyssar/msdcli/utils"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/lyssar/msdcli/config"

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

func DownloadInstaller() (utils.MinecraftVersion, ForgeVersion) {
	var minecraftVersion utils.MinecraftVersion
	var forgeVersion ForgeVersion
	if *config.McVersion == "" {
		minecraftVersion = utils.SelectMinecraftVersion()
	} else {
		minecraftVersion = utils.GetMinecraftVersionInfo(*config.McVersion)
	}

	fmt.Printf("MC Version: %s\n", minecraftVersion.ID)

	if *config.ServerVersion == "" {
		forgeVersion = selectForgeVersion(minecraftVersion)
	} else {
		forgeVersion = ForgeVersion{Version: *config.ServerVersion, Installer: getMavenDownloadLink(minecraftVersion.ID, *config.ServerVersion)}
	}

	fmt.Printf("Forge Version: %s\n", forgeVersion.Version)

	downloadForge(forgeVersion)

	eulaFile := []byte("forge-" + minecraftVersion.ID + "-" + forgeVersion.Version)
	err := os.WriteFile("version.txt", eulaFile, 0644)

	if err != nil {
		log.Fatal(err)
	}

	return minecraftVersion, forgeVersion
}

func InstalServer(minecraftVersion utils.MinecraftVersion, forgeVersion ForgeVersion) {
	var err error
	utils.CreateEula()

	out, cmdErr := exec.Command("/usr/bin/java", "-jar", utils.GetCwd()+"/"+config.InstallerFile, "--installServer").Output()

	if cmdErr != nil {
		log.Fatal(cmdErr)
	}
	fmt.Println(string(out))

	renamErr := os.Rename("forge-"+minecraftVersion.ID+"-"+forgeVersion.Version+".jar", config.ServerFile)
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

func selectForgeVersion(minecraftVersion utils.MinecraftVersion) ForgeVersion {
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

func getForgeVersionForMinecraftVersion(mcVersion utils.MinecraftVersion) []ForgeVersion {
	versionList := []ForgeVersion{}

	fmt.Printf("Loading forge version list for minecraft %s.\n", mcVersion.ID)
	loadUrl := fmt.Sprintf("https://files.minecraftforge.net/net/minecraftforge/forge/index_%s.html", mcVersion.ID)
	doc, err := htmlquery.LoadURL(loadUrl)
	if err != nil {
		log.Fatal(err)
	}

	links := htmlquery.Find(doc, "//table[contains(@class, 'download-list')]/tbody/tr/td[contains(@class, 'download-version')]")
	for _, n := range links {
		forgeVersionNumberString := strings.Trim(htmlquery.InnerText(n), " \n")
		installer := getMavenDownloadLink(mcVersion.ID, forgeVersionNumberString)

		forgeVersion := ForgeVersion{Version: forgeVersionNumberString, Installer: installer}

		versionList = append(versionList, forgeVersion)
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
	out, err := os.Create(utils.GetCwd() + "/" + config.InstallerFile)
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

func getMavenDownloadLink(mcVersion string, forgeVersion string) string {
	return fmt.Sprintf("https://maven.minecraftforge.net/net/minecraftforge/forge/%[1]s-%[2]s/forge-%[1]s-%[2]s-installer.jar", mcVersion, forgeVersion)
}
