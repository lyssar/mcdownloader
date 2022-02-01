package config

import (
	"flag"
	"log"
	"os"
)

var (
	ServerType          *string
	McVersion           *string
	ServerVersion       *string
	PackageId           *int
	ServerPackageFileID *int
)

func LoadArgs(subcommand string) {
	var modpackFlags *flag.FlagSet
	switch subcommand {
	case "server":
		modpackFlags = flag.NewFlagSet("server", flag.ExitOnError)
		ServerType = modpackFlags.String("type", "", "Server type. [vanilla|fabric|forge|papermc|spigot]")
		McVersion = modpackFlags.String("mcversion", "", "Minecraft version.")
		ServerVersion = modpackFlags.String("serverVersion", "", "Server version eq. version of your server type.")
	case "modpack":
		modpackFlags = flag.NewFlagSet("modpack", flag.ExitOnError)
		PackageId = modpackFlags.Int("packageId", 0, "Modpack ID from curseforge")
		ServerPackageFileID = modpackFlags.Int("serverPackageFileID", 0, "File ID to download (excplicit server version of a package)")
	}

	if modpackFlags != nil {
		err := modpackFlags.Parse(os.Args[2:])

		if err != nil {
			log.Fatal(err)
		}
	}
}
