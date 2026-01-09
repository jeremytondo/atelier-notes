package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

// Config holds the application configuration
type Config struct {
	NotesDir string
}

// Load initializes and loads the configuration
func Load() {
	// 1. Check XDG_CONFIG_HOME
	if xdgConfigHome := os.Getenv("XDG_CONFIG_HOME"); xdgConfigHome != "" {
		viper.AddConfigPath(filepath.Join(xdgConfigHome, "atelier-notes"))
	}

	// 2. Check ~/.config/atelier-notes (Standard XDG fallback, preferred by many macOS devs)
	if home, err := os.UserHomeDir(); err == nil {
		viper.AddConfigPath(filepath.Join(home, ".config", "atelier-notes"))
	}

	// 3. Fallback to OS default (e.g., ~/Library/Application Support on macOS)
	if configDir, err := os.UserConfigDir(); err == nil {
		viper.AddConfigPath(filepath.Join(configDir, "atelier-notes"))
	}

	viper.SetConfigName("settings")
	viper.SetConfigType("yaml")

	viper.SetEnvPrefix("ATELIER_NOTES")
	viper.AutomaticEnv()

	// Define defaults
	viper.SetDefault("notes-dir", ".")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			fmt.Println("Error reading config file:", err)
		}
	}
}

// GetNotesDir returns the configured notes directory
func GetNotesDir() string {
	path := viper.GetString("notes-dir")
	if strings.HasPrefix(path, "~/") {
		home, err := os.UserHomeDir()
		if err == nil {
			return filepath.Join(home, path[2:])
		}
	}
	return path
}
