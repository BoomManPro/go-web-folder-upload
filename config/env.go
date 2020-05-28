package config

import "go.uber.org/config"

type StoreConfig struct {
	StorePath string `yaml:"store-path"`
}

func GetApplicationConfigFromYml(path string) *StoreConfig {

	fileYml := config.File(path)

	provider, err := config.NewYAML(fileYml)
	if err != nil {
		panic(err) // handle error
	}
	var storeConfig StoreConfig
	if err := provider.Get("application.web-folder-upload").Populate(&storeConfig); err != nil {
		panic(err)
	}

	return &storeConfig
}
