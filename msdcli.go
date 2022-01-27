package main

import (
	"fmt"
	"os"

	"github.com/lyssar/msdcli/config"
	"github.com/lyssar/msdcli/modpacks"
	"github.com/lyssar/msdcli/server"
)

func main() {
	subcommand := ""

	if len(os.Args) > 1 {
		subcommand = os.Args[1]
	}

	config.LoadArgs(subcommand)

	switch subcommand {
	case "server":
		server.InstalServer()
	case "modpack":
		modpacks.Download()
	default:
		fmt.Printf("msdcli <server|modpack> ...args:\n")
		fmt.Printf("  %s\n    \tDisplay this help\n", subcommand)
		fmt.Printf("  server\n")
		fmt.Printf("      \tInstall a minecraft server with. Available are [fabric|forge|papermc|spigot] as type.\n")
		fmt.Printf("  modpack\n")
		fmt.Printf("      \tDownload curseforge modpack either by given args or with tty\n")

	}
}
