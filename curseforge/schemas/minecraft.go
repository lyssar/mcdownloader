package schemas

import "time"

type ModLoaderInstallMethod int

const (
	ModLoaderInstallMethodForgeInstaller   ModLoaderInstallMethod = 1
	ModLoaderInstallMethodForgeJarInstall  ModLoaderInstallMethod = 2
	ModLoaderInstallMethodForgeInstallerV2 ModLoaderInstallMethod = 3
)

type ModLoaderType int

const (
	ModLoaderTypeAny        FileStatus = 0
	ModLoaderTypeForge      FileStatus = 1
	ModLoaderTypeCauldron   FileStatus = 2
	ModLoaderTypeLiteLoader FileStatus = 3
	ModLoaderTypeFabric     FileStatus = 4
	ModLoaderTypeQuilt      FileStatus = 5
)

type ApiResponseOfListOfMinecraftGameVersion struct {
	Data []MinecraftGameVersion `json:"data"`
}

type ApiResponseOfListOfMinecraftModLoaderIndex struct {
	Data []MinecraftModLoaderIndex `json:"data"`
}

type ApiResponseOfMinecraftGameVersion struct {
	Data MinecraftGameVersion `json:"data"`
}

type ApiResponseOfMinecraftModLoaderVersion struct {
	Data MinecraftModLoaderVersion `json:"data"`
}

type MinecraftGameVersion struct {
	ID                    int                   `json:"id"`
	GameVersionID         int                   `json:"gameVersionId"`
	VersionString         string                `json:"versionString"`
	JarDownloadURL        string                `json:"jarDownloadUrl"`
	JSONDownloadURL       string                `json:"jsonDownloadUrl"`
	Approved              bool                  `json:"approved"`
	DateModified          time.Time             `json:"dateModified"`
	GameVersionTypeID     int                   `json:"gameVersionTypeId"`
	GameVersionStatus     GameVersionStatus     `json:"gameVersionStatus"`
	GameVersionTypeStatus GameVersionTypeStatus `json:"gameVersionTypeStatus"`
}

type MinecraftModLoaderIndex struct {
	Name         string    `json:"name"`
	GameVersion  string    `json:"gameVersion"`
	Latest       bool      `json:"latest"`
	Recommended  bool      `json:"recommended"`
	DateModified time.Time `json:"dateModified"`
	Type         int       `json:"type"`
}

type MinecraftModLoaderVersion struct {
	ID                             int                    `json:"id"`
	GameVersionID                  int                    `json:"gameVersionId"`
	MinecraftGameVersionID         int                    `json:"minecraftGameVersionId"`
	ForgeVersion                   string                 `json:"forgeVersion"`
	Name                           string                 `json:"name"`
	Type                           ModLoaderType          `json:"type"`
	DownloadURL                    string                 `json:"downloadUrl"`
	Filename                       string                 `json:"filename"`
	InstallMethod                  ModLoaderInstallMethod `json:"installMethod"`
	Latest                         bool                   `json:"latest"`
	Recommended                    bool                   `json:"recommended"`
	Approved                       bool                   `json:"approved"`
	DateModified                   time.Time              `json:"dateModified"`
	MavenVersionString             string                 `json:"mavenVersionString"`
	VersionJSON                    string                 `json:"versionJson"`
	LibrariesInstallLocation       string                 `json:"librariesInstallLocation"`
	MinecraftVersion               string                 `json:"minecraftVersion"`
	AdditionalFilesJSON            string                 `json:"additionalFilesJson"`
	ModLoaderGameVersionID         int                    `json:"modLoaderGameVersionId"`
	ModLoaderGameVersionTypeID     int                    `json:"modLoaderGameVersionTypeId"`
	ModLoaderGameVersionStatus     GameVersionStatus      `json:"modLoaderGameVersionStatus"`
	ModLoaderGameVersionTypeStatus GameVersionTypeStatus  `json:"modLoaderGameVersionTypeStatus"`
	McGameVersionID                int                    `json:"mcGameVersionId"`
	McGameVersionTypeID            int                    `json:"mcGameVersionTypeId"`
	McGameVersionStatus            GameVersionStatus      `json:"mcGameVersionStatus"`
	McGameVersionTypeStatus        GameVersionTypeStatus  `json:"mcGameVersionTypeStatus"`
	InstallProfileJSON             string                 `json:"installProfileJson"`
}
