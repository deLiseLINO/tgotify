package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

// Config represents the configuration structure.
type Config struct {
	TgHost  string        `mapstructure:"tg_host"`
	Storage StorageConfig `mapstructure:"storage"`
}

// StorageConfig represents the database storage configuration.
type StorageConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Database string `mapstructure:"database"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

// Init initializes the configuration by reading values from both environment variables and a YAML file.
func Init() *Config {
	cfg := &Config{}
	cfg.ReadEnv()
	cfg.ReadYaml()
	return cfg
}

// ReadEnv reads environment variables into the configuration using Viper.
func (*Config) ReadEnv() {
	// Set the configuration file to .env and automatically read environment variables.
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	// Attempt to read the configuration from the .env file.
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("failed to read .env: %v", err)
	}
}

// ReadYaml reads configuration values from a YAML file specified in the "CONFIG_PATH" environment variable.
func (cfg *Config) ReadYaml() {
	// Retrieve the path to the YAML configuration file from the "CONFIG_PATH" environment variable.
	configPath := viper.GetString("CONFIG_PATH")

	// Check if the configPath is empty.
	if configPath == "" {
		log.Fatalf("CONFIG_PATH environment variable is not set or is empty")
	}

	// Check if the file exists; if not, log an error and exit.
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file path: %s does not exist: %v", configPath, err)
	}

	// Set the configuration file to the specified YAML file and configure it as YAML format.
	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")

	// Read the configuration from the specified YAML file.
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("failed to read the config file: %v", err)
	}

	// Unmarshal the configuration into the provided Config structure.
	if err := viper.Unmarshal(cfg); err != nil {
		log.Fatalf("failed to unmarshal config: %v", err)
	}
}
