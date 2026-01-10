package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/viper"
)

func TestGetNotesDir(t *testing.T) {
	viper.Reset()
	defer viper.Reset()

	// Test default
	viper.SetDefault("notes-dir", ".")
	if GetNotesDir() != "." {
		t.Errorf("Expected '.', got %s", GetNotesDir())
	}

	// Test absolute path
	absPath := "/tmp/notes"
	viper.Set("notes-dir", absPath)
	if GetNotesDir() != absPath {
		t.Errorf("Expected %s, got %s", absPath, GetNotesDir())
	}

	// Test tilde expansion
	home, _ := os.UserHomeDir()
	viper.Set("notes-dir", "~/notes")
	expected := filepath.Join(home, "notes")
	if GetNotesDir() != expected {
		t.Errorf("Expected %s, got %s", expected, GetNotesDir())
	}
}
