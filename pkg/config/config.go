package config

import (
	"bytes"
	"embed"
	"fmt"

	"github.com/spf13/viper"
)

//go:embed config.json
var configFile embed.FS //embeding config.json with this variable

func Init() {

	data, err := configFile.ReadFile("config.json")
	if err != nil {
		fmt.Printf("Error reading config file: %v\n", err)
		return
	}

	viper.SetConfigType("json")

	if err := viper.ReadConfig(bytes.NewBuffer(data)); err != nil {
		fmt.Printf("Error reading config: %v\n", err)
		panic("Error reading config: " + err.Error())

	}

}

func GetFolderStructure(folder_name string) []string {
	arrayValues := viper.GetStringSlice(folder_name)
	return arrayValues
}
