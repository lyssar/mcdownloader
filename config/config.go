package config

import (
	"bytes"
	_ "embed"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Config struct {
	CurseForge struct {
		BaseUrlProtocol string `yaml:"baseUrlProtocol"`
		BaseUrl         string `yaml:"baseUrl"`
		ApiKey          string `yaml:"apiKey"`
		MinecraftGameID int    `yaml:"minecraftGameID"`
	} `yaml:"curseforge"`
	MinecraftVersion string          `yaml:"minecraftVersion"`
	ForgeVersion     string          `yaml:"forgeVersion"`
	Minecraft        MinecraftConfig `yaml:"minecraft"`
}

type MinecraftConfig struct {
	MetaJson string `yaml:"metajson"`
}

var loadedConfig *Config

//go:embed resource/api.yaml
var apiResource []byte

func GetConfig() *Config {
	if loadedConfig == nil {
		_, err := LoadConfig()
		cobra.CheckErr(err)
	}
	return loadedConfig
}

func LoadConfig() (config *Config, err error) {
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	err = viper.ReadConfig(bytes.NewBuffer(apiResource))
	cobra.CheckErr(err)

	err = viper.Unmarshal(&config)
	loadedConfig = config
	return
}
