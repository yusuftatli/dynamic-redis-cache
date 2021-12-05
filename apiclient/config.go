package apiclient

import (
	"fmt"
	"os"
	"path"

	"gopkg.in/yaml.v2"
)

type ProviderConfig struct {
	Provider1 Provider1Config `yaml:"provider1"`
	Provider2 Provider2Config `yaml:"provider2"`
	Provider3 Provider3Config `yaml:"provider3"`
}

// type ProviderConfig struct {
// 	Provider1 struct {
// 		url string `yaml:"url"`
// 	} `yaml:"provider1"`
// 	Provider2 struct {
// 		url string `yaml:"url"`
// 	} `yaml:"provider2"`
// }

func NewDefaultConfig() (*ProviderConfig, error) {
	configFilePath, err := getFilePath("config.yml")
	if err != nil {
		return nil, err
	}

	file, err := os.Open(configFilePath)
	if err != nil {
		fmt.Errorf("Failed to open config.yml file", err)
		// errors.Wrap(err, "Failed to open config.yml file")
		return nil, err
	}

	defer file.Close()

	cfg := new(ProviderConfig)
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&cfg)
	if err != nil {
		fmt.Errorf("Failed to decode yaml config", err)
	}

	return cfg, nil
}

func getFilePath(fileName string) (string, error) {
	rootPath, err := os.Getwd()
	if err != nil {
		// errors.Wrap(err, "Failed to get root path")
		return "", err
	}

	return path.Join(path.Dir(rootPath), "/hepsiburada/"+fileName), nil
}
