package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

const defaultConfigPath = "./configs"

func LoadConfig(configPath, name string) (*SatoriConfig, error) {
	if configPath == "" {
		configPath = defaultConfigPath
	}
	exists := pathExists(configPath)
	if !exists {
		return nil, fmt.Errorf("config path not exists: %v", configPath)
	}
	c := viper.New()
	c.AddConfigPath(configPath)
	c.SetConfigName(name)
	c.SetConfigType("yaml")
	err := c.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	conf := &SatoriConfig{}
	err = c.Unmarshal(conf)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return conf, nil
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
