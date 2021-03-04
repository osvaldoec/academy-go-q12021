package config

import (
	"github.com/spf13/viper"
)

// Configuration - struct to handle variables
type Configuration struct {
	Address   string
	PokemonDB string
}

// ReadConfig - handles main config
func ReadConfig(configFile string) (*Configuration, error) {
	public := viper.New()
	public.SetConfigFile(configFile)
	if err := public.ReadInConfig(); err != nil {
		return nil, err
	}
	config := &Configuration{}

	err := public.Unmarshal(config)
	if err != nil {
		return nil, err
	}

	public.WatchConfig()

	return config, nil
}
