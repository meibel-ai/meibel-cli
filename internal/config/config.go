package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

// Init initializes the configuration.
// It reads from the config file if provided, otherwise looks in default locations.
func Init(cfgFile string) error {
	if cfgFile != "" {
		// Use config file from the flag
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory
		home, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("could not find home directory: %w", err)
		}

		// Search config in $HOME/.appname
		configDir := filepath.Join(home, ".meibel")
		viper.AddConfigPath(configDir)
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")
	}

	// Read environment variables
	viper.SetEnvPrefix("meibel")
	viper.AutomaticEnv()

	// Read in config file if it exists
	if err := viper.ReadInConfig(); err != nil {
		// It's okay if config file doesn't exist
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return fmt.Errorf("error reading config file: %w", err)
		}
	}

	return nil
}

// GetString returns a string config value.
func GetString(key string) string {
	return viper.GetString(key)
}

// GetInt returns an int config value.
func GetInt(key string) int {
	return viper.GetInt(key)
}

// GetBool returns a bool config value.
func GetBool(key string) bool {
	return viper.GetBool(key)
}

// Set sets a config value.
func Set(key string, value interface{}) {
	viper.Set(key, value)
}
