package main

import (
	"flag"
	"fmt"
	"github.com/lyssar/mcdownloader/server"
	"os"
)

func main() {
	subcommand := ""

	modpackFlags := flag.NewFlagSet("modpack", flag.ExitOnError)
	packageId := modpackFlags.String("packageId", "", "packageId")
	packageVersion := modpackFlags.String("version", "", "version")

	if len(os.Args) > 1 {
		subcommand = os.Args[1]
	}

	switch subcommand {
	case "server":
		server.InstalServer()
	case "modpack":
		modpackFlags.Parse(os.Args[2:])
		fmt.Println("Start Installing MC Modpack to Server")
		fmt.Println("    package: ", *packageId)
		fmt.Println("    version: ", *packageVersion)
	default:
		fmt.Println("HELP")
	}
}
