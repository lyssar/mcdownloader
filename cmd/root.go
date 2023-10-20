package cmd

import (
	"github.com/lyssar/msdcli/config"
	"github.com/spf13/viper"
	"os"

	cc "github.com/ivanpirog/coloredcobra"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "msdcli",
	Short: "",
	Long: `CLI to create server instances with and without mods or modpacks.
With the possibility to do this in an TTY less context through adding parameters.
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cc.Init(&cc.Config{
		RootCmd:         rootCmd,
		Headings:        cc.HiCyan + cc.Bold + cc.Underline,
		Commands:        cc.HiYellow + cc.Bold,
		Example:         cc.Italic,
		ExecName:        cc.Bold,
		Flags:           cc.HiYellow + cc.Bold,
		NoExtraNewlines: true,
	})

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().BoolP("debug", "d", false, "Forcing the log level to debug")
	rootCmd.PersistentFlags().StringP("working-dir", "w", "", "Folder to execute creation in")
	_ = viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))
	_ = viper.BindPFlag("working-dir", rootCmd.PersistentFlags().Lookup("working-dir"))
}

func initConfig() {
	_, err := config.LoadConfig()
	cobra.CheckErr(err)
}
