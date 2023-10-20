package forge

import (
	"fmt"
	forgeVersionApi "github.com/kleister/go-forge/version"
	"github.com/lyssar/msdcli/errors"
	"github.com/lyssar/msdcli/logger"
	"github.com/lyssar/msdcli/utils"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"sort"
	"strings"
)

type ForgeApi struct {
	forgeVersionList     forgeVersionApi.Response
	SelectedForgeVersion *forgeVersionApi.Version
	MinecraftVersion     string
}

func NewForgeClient() *ForgeApi {
	forgeResultSet, err := forgeVersionApi.FromDefault()
	cobra.CheckErr(err)
	return &ForgeApi{
		forgeVersionList: forgeResultSet,
	}
}

func (api *ForgeApi) verifyForgeVersion(forgeVersion forgeVersionApi.Version, minecraftVersionString string) (*bool, error) {
	ok := true

	if forgeVersion.Minecraft != minecraftVersionString {
		err := errors.NewError(
			fmt.Sprintf(
				"Forge version %s is not for minecraft %s. Expecting minecraft version %s",
				forgeVersion.ID,
				minecraftVersionString,
				forgeVersion.Minecraft,
			),
		)

		return nil, err
	}

	return &ok, nil
}

func (api *ForgeApi) getFilteredReleasesForMinecraftVersion(minecraftVersionString string) forgeVersionApi.Versions {
	forgeMinecraftVersionFilter := &forgeVersionApi.Filter{
		Minecraft: minecraftVersionString,
	}
	logger.Debugf("Used version filter: %#v", forgeMinecraftVersionFilter)
	forgeVersionList := api.forgeVersionList.Releases.Filter(forgeMinecraftVersionFilter)

	return api.reverseSortVersionList(forgeVersionList)
}

func (api *ForgeApi) getSpecificForgeVersion(forgeVersionString string) forgeVersionApi.Versions {
	forgeMinecraftVersionFilter := &forgeVersionApi.Filter{
		Version: forgeVersionString,
	}

	forgeVersionList := api.forgeVersionList.Releases.Filter(forgeMinecraftVersionFilter)

	return api.reverseSortVersionList(forgeVersionList)
}

func (api *ForgeApi) reverseSortVersionList(forgeVersionList forgeVersionApi.Versions) forgeVersionApi.Versions {
	sort.Slice(forgeVersionList, func(i, j int) bool {
		return forgeVersionList[i].ID > forgeVersionList[j].ID
	})

	return forgeVersionList
}

func (api *ForgeApi) SelectForgeVersion(forgeVersionString string, minecraftVersionString string) (bool, *errors.ApplicationError) {
	logger.Info("Choosing forge version")

	if forgeVersionString == "" {
		forgeVersionList := api.getFilteredReleasesForMinecraftVersion(minecraftVersionString)
		selectedForgeVersion, err := RenderSelect(forgeVersionList)
		if err != nil {
			return false, errors.NewError(err.Error())
		}

		api.SelectedForgeVersion = selectedForgeVersion

		return true, nil
	} else {
		forgeVersionList := api.getSpecificForgeVersion(forgeVersionString)

		if len(forgeVersionList) == 0 {
			forgeError := errors.NewError(fmt.Sprintf("No forge package found for %s", forgeVersionString))
			return false, forgeError
		}
		_, err := api.verifyForgeVersion(forgeVersionList[0], minecraftVersionString)

		if err != nil {
			return false, errors.NewError(err.Error())
		}
		api.SelectedForgeVersion = &forgeVersionList[0]
		return true, nil

	}
}

func (api *ForgeApi) DownloadServer(workingDir string) (bool, *errors.ApplicationError) {
	if api.SelectedForgeVersion == nil {
		return false, errors.NewError("Forge version not selected.")
	}
	downloadFile := strings.Replace(api.SelectedForgeVersion.URL, "-universal", "-installer", -1)
	logger.Infof("Download forge server for version %s", api.SelectedForgeVersion.Minecraft)
	logger.Debugf("Used forge download file: %s", downloadFile)

	if _, err := os.Stat(workingDir); os.IsNotExist(err) {
		errors.CheckStandardErr(err)
	}
	utils.DownloadFile(downloadFile, "forge.jar", workingDir)
	return true, nil
}

func (api *ForgeApi) InstallServer(workingDir string) (bool, *errors.ApplicationError) {
	logger.Info("Installing forge version")
	if _, err := os.Stat(workingDir); os.IsNotExist(err) {
		errors.CheckStandardErr(err)
	}

	forgerInstallerPath := fmt.Sprintf("%s/forge.jar", workingDir)
	_, err := os.Stat(forgerInstallerPath)

	if os.IsNotExist(err) {
		errors.CheckStandardErr(err)
	}
	logger.Debugf("Found forge.jar. Trying to install it now.")

	cmd := exec.Command("java", "-jar", forgerInstallerPath, "--installServer")
	cmd.Dir = workingDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	errors.CheckStandardErr(err)

	err = os.Remove(forgerInstallerPath)
	errors.CheckStandardErr(err)

	return true, nil
}
