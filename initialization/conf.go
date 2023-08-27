package initialization

import (
	"blog_server/config"
	"blog_server/global"
	"io/fs"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

const ConfigFile = "settings.yaml"

func InitConf() {
	c := &config.Config{}
	yamlConf, err := os.ReadFile(ConfigFile)
	if err != nil {
		log.Fatalf("Get yamlConf error: %s\n", err)
	}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("Get yaml Unmarchsal error: %s\n", err)
	}
	log.Println("Configuration file loads successfully.")
	global.Config = c
}

func SettingYaml() error {
	byteFile, err := yaml.Marshal(global.Config)

	if err != nil {
		return err
	}

	err = os.WriteFile(ConfigFile, byteFile, fs.ModePerm)

	if err != nil {
		return err
	}

	return nil
}
