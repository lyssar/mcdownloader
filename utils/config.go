package utils

import (
	"bytes"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
)

type Config struct {
	CurseForge struct {
		BaseUrlProtocol string `yaml:"baseUrlProtocol"`
		BaseUrl         string `yaml:"baseUrl"`
		ApiKey          string `yaml:"apiKey"`
		MinecraftGameID int    `yaml:"minecraftGameID"`
	} `yaml:"curseforge"`
	MinecraftVersion string `yaml:"minecraftVersion"`
	ForgeVersion     string `yaml:"forgeVersion"`
	Minecraft        struct {
		MetaJson string `yaml:"metajson"`
	} `yaml:"minecraft"`
}

var loadedConfig *Config

func GetConfig() *Config {
	if loadedConfig == nil {
		_, err := LoadConfig()
		cobra.CheckErr(err)
	}
	return loadedConfig
}

func LoadConfig() (config *Config, err error) {
	var configData, _ = ioutil.ReadFile("resource/api.yaml")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	err = viper.ReadConfig(bytes.NewBuffer(configData))
	cobra.CheckErr(err)

	err = viper.Unmarshal(&config)
	loadedConfig = config
	return
}
