package forge

import (
	"errors"
	"fmt"
	"github.com/gookit/color"
	"github.com/hashicorp/go-version"
	forgeVersionApi "github.com/kleister/go-forge/version"
	"github.com/lyssar/msdcli/minecraft"
	"github.com/lyssar/msdcli/utils"
	"github.com/spf13/cobra"
	"os/exec"
	"regexp"
	"strconv"
)

func CreateServer(config *utils.Config) {
	utils.PrintInfo("Creating forge server.")
	minecraftMetaApi := minecraft.NewMinecraftMetaApi(config.Minecraft.MetaJson)
	selectedMinecraftVersion, err := minecraftMetaApi.FindMinecraftVersion(config.MinecraftVersion)
	cobra.CheckErr(err)

	forgeApp := NewForgeClient()

	var selectedForgeVersion *forgeVersionApi.Version
	selectedForgeVersion, err = forgeApp.SelectForgeVersion(config.ForgeVersion, selectedMinecraftVersion.ID)
	cobra.CheckErr(err)

	fmt.Println(selectedForgeVersion.URL)

	mcRelease := selectedMinecraftVersion.DownloadRelease()
	_, err = checkJavaVersion(mcRelease)
	cobra.CheckErr(err)

	// Download server jar
	if mcRelease.DownloadServer() {
		// Run and setup once
		mcRelease.InstallServer()
	}

	fmt.Println("== FORGE ===============================================")
	fmt.Println(mcRelease.Downloads.Server.URL)
	// Download forge
	// Run and setup once
}

func checkJavaVersion(mcRelease minecraft.McRelease) (bool, error) {
	cmdPrep := "java --version"
	cmdOutput, _ := exec.Command("bash", "-c", cmdPrep).CombinedOutput()

	utils.PrintInfo("Check host java version.")
	versionRegex := regexp.MustCompile(`^.*\s(\d+).\d+\.\d+`)
	completeVersionStringRegex := regexp.MustCompile(`^.*\s(\d+.\d+\.\d+)`)
	matches := versionRegex.FindStringSubmatch(string(cmdOutput))
	completeStringMatches := completeVersionStringRegex.FindStringSubmatch(string(cmdOutput))
	if len(matches) > 0 {
		utils.PrintInfo("Host java version found.")
		javaVersion := string(matches[1])
		completeJavaVersion := string(completeStringMatches[1])
		majorVersion := strconv.Itoa(mcRelease.JavaVersion.MajorVersion)

		releaseMajorVersion, err := version.NewVersion(majorVersion)
		cobra.CheckErr(err)

		hostJavaVersion, err := version.NewVersion(javaVersion)
		cobra.CheckErr(err)

		if !hostJavaVersion.Equal(releaseMajorVersion) {
			return false, errors.New(color.Error.Sprintf("Java version %s found, expected %s.x.x. Please install the correct java version", completeJavaVersion, majorVersion))
		}
	} else {
		majorVersion := strconv.Itoa(mcRelease.JavaVersion.MajorVersion)
		return false, errors.New(color.Error.Sprintf("Couldn't find any java version. Please install java %s", majorVersion))
	}

	return true, nil
}
