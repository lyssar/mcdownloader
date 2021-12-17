package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/lyssar/mcdownloader/server"
	"github.com/manifoldco/promptui"
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
		fmt.Println("Start Installing MC Server")
		prompt := promptui.Select{
			Label: "Select server type",
			Items: []string{"Forge", "Fabric", "Spitgot", "PaperMC"},
		}
		_, result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
		fmt.Printf("You choose %q\n", result)

		server.Forge()
	case "modpack":
		modpackFlags.Parse(os.Args[2:])
		fmt.Println("Start Installing MC Modpack to Server")
		fmt.Println("    package: ", *packageId)
		fmt.Println("    version: ", *packageVersion)
	default:
		fmt.Println("HELP")
	}
}
