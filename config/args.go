package config

import (
	"flag"
	"log"
	"os"
)

var (
	ServerType     *string
	McVersion      *string
	ServerVersion  *string
	PackageId      *int
	PackageVersion *string
)

func LoadArgs(subcommand string) {
	var modpackFlags *flag.FlagSet
	switch subcommand {
	case "server":
		modpackFlags = flag.NewFlagSet("server", flag.ExitOnError)
		ServerType = modpackFlags.String("type", "", "type")
		McVersion = modpackFlags.String("mcversion", "", "mcversion")
		ServerVersion = modpackFlags.String("serverVersion", "", "serverVersion")
	case "modpack":
		modpackFlags = flag.NewFlagSet("modpack", flag.ExitOnError)
		PackageId = modpackFlags.Int("packageId", 0, "packageId")
		PackageVersion = modpackFlags.String("version", "", "version")
	}
	if modpackFlags != nil {
		err := modpackFlags.Parse(os.Args[2:])

		if err != nil {
			log.Fatal(err)
		}
	}
}
