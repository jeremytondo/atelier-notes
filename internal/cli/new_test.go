package cli

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestNewCommand(t *testing.T) {
	// Setup temp dir for notes
	tempDir, err := os.MkdirTemp("", "cli-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Set targetDir to tempDir to override config
	targetDir = tempDir

	// Buffer to capture stdout
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	// Execute command: atelier-notes new "My Test Note"
	rootCmd.SetArgs([]string{"new", "My Test Note"})
	err = rootCmd.Execute()
	if err != nil {
		t.Fatalf("Command execution failed: %v", err)
	}

	// Verify output contains the path
	output := strings.TrimSpace(buf.String())
	if !strings.HasSuffix(output, "my-test-note.md") {
		t.Errorf("Expected output to end with my-test-note.md, got %s", output)
	}

	// Verify file was created
	if _, err := os.Stat(output); os.IsNotExist(err) {
		t.Errorf("File was not created at %s", output)
	}
}
