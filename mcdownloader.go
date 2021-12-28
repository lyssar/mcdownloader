package main

import (
	"fmt"
	"github.com/lyssar/mcdownloader/config"
	"github.com/lyssar/mcdownloader/server"
	"os"
)

func main() {
	subcommand := ""

	if len(os.Args) > 1 {
		subcommand = os.Args[1]
	}

	config.LoadArgs(subcommand)

	switch subcommand {
	case "server":
		fmt.Println("    McVersion: ", *config.McVersion)
		fmt.Println("    ServerType: ", *config.ServerType)
		fmt.Println("    ServerVersion: ", *config.ServerVersion)
		server.InstalServer()
	case "modpack":
		config.LoadArgs("server")
		fmt.Println("Start Installing MC Modpack to Server")
		fmt.Println("    package: ", *config.PackageId)
		fmt.Println("    version: ", *config.PackageVersion)
	default:
		fmt.Println("HELP")
	}
}
