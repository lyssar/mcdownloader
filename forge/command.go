package forge

import (
	"fmt"
	"github.com/hashicorp/go-version"
	"github.com/kyokomi/emoji/v2"
	"github.com/lyssar/msdcli/config"
	"github.com/lyssar/msdcli/errors"
	"github.com/lyssar/msdcli/logger"
	"github.com/lyssar/msdcli/minecraft"
	"os"
	"os/exec"
	"regexp"
	"strconv"
)

func CreateServer(workingDir string, config *config.Config) {
	_, _ = emoji.Println(":compass:")
	logger.Infof("%s Start creating forge server.", emoji.Emojize(":genie:"))
	_, err := checkProjectDir(workingDir)
	errors.Check(err)

	minecraftMetaApi := minecraft.NewMinecraftMetaApi(config.Minecraft.MetaJson)
	selectedMinecraftVersion, err := minecraftMetaApi.FindMinecraftVersion(config.MinecraftVersion)
	errors.Check(err)

	forgeApp := NewForgeClient()

	_, err = forgeApp.SelectForgeVersion(config.ForgeVersion, selectedMinecraftVersion.ID)
	errors.Check(err)

	mcRelease := selectedMinecraftVersion.DownloadRelease()
	_, err = checkJavaVersion(mcRelease)
	errors.Check(err)

	// Download server jar
	if mcRelease.DownloadServer(workingDir) {
		logger.Debug("Start server to install once")
		// Run and setup once
		mcRelease.InstallServer(workingDir)
	}

	_, err = forgeApp.DownloadServer(workingDir)
	errors.Check(err)

	_, err = forgeApp.InstallServer(workingDir)
	errors.Check(err)

	logger.Infof("%s Forge server successfully installed. Have fun playing.", emoji.Sprint(":partying_face:"))
	logger.Infof("To run your instance execute %[1]s/run.sh nogui or %[1]s/run.bat nogui.", workingDir)
}

func checkProjectDir(folder string) (bool, *errors.ApplicationError) {
	err := os.MkdirAll(folder, 0755)
	if err != nil {
		return false, errors.NewError("Couldn't create directory.")
	}

	f, err := os.Open(folder)
	errors.CheckStandardErr(err)

	defer f.Close()

	_, err = f.Readdirnames(1)
	if err == nil {
		return false, errors.NewError("Folder is not empty. Can't create forge server.")
	}
	return true, nil
}

func checkJavaVersion(mcRelease minecraft.McRelease) (bool, *errors.ApplicationError) {
	cmdPrep := "java --version"
	cmdOutput, _ := exec.Command("bash", "-c", cmdPrep).CombinedOutput()

	logger.Info("Checking host java version.")
	logger.Debug("Using java --version")
	versionRegex := regexp.MustCompile(`^.*\s(\d+).\d+\.\d+`)
	completeVersionStringRegex := regexp.MustCompile(`^.*\s(\d+.\d+\.\d+)`)
	matches := versionRegex.FindStringSubmatch(string(cmdOutput))
	completeStringMatches := completeVersionStringRegex.FindStringSubmatch(string(cmdOutput))

	if len(matches) > 0 {
		logger.Debug("Host java version found.")
		javaVersion := string(matches[1])
		completeJavaVersion := string(completeStringMatches[1])
		majorVersion := strconv.Itoa(mcRelease.JavaVersion.MajorVersion)

		releaseMajorVersion, err := version.NewVersion(majorVersion)
		if err != nil {
			return false, errors.NewError(err.Error())
		}

		hostJavaVersion, err := version.NewVersion(javaVersion)
		if err != nil {
			return false, errors.NewError(err.Error())
		}

		if !hostJavaVersion.Equal(releaseMajorVersion) {
			err := errors.NewWarning(fmt.Sprintf("Java version %s found, expected %s.x.x. Please install the correct java version", completeJavaVersion, majorVersion))
			return false, err
		}
	} else {
		logger.Debug(fmt.Sprintf("Return of check command: %s", string(cmdOutput)))
		majorVersion := strconv.Itoa(mcRelease.JavaVersion.MajorVersion)
		err := errors.NewFatal(fmt.Sprintf("Couldn't find any java version. Please install java %s", majorVersion))
		return false, err
	}

	return true, nil
}
