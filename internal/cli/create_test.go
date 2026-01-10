package cli

import (
	"bytes"
	"os"
	"strings"
	"testing"
	"time"
)

func TestCreateCommand(t *testing.T) {
	// Setup temp dir for notes
	tempDir, err := os.MkdirTemp("", "cli-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer func() { _ = os.RemoveAll(tempDir) }()

	// Set targetDir to tempDir to override config
	targetDir = tempDir

	// Buffer to capture stdout
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	// Execute command: atelier-notes create "My Test Note"
	rootCmd.SetArgs([]string{"create", "My Test Note"})
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

func TestCreateDailyCommand(t *testing.T) {
	// Setup temp dir for notes
	tempDir, err := os.MkdirTemp("", "cli-test-daily")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer func() { _ = os.RemoveAll(tempDir) }()

	// Set targetDir to tempDir
	targetDir = tempDir

	// Buffer to capture stdout
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	// Execute command: atelier-notes create --daily
	rootCmd.SetArgs([]string{"create", "--daily"})
	err = rootCmd.Execute()
	if err != nil {
		t.Fatalf("Command execution failed: %v", err)
	}

	output := strings.TrimSpace(buf.String())

	// Expect filename like daily-20260110.md
	today := time.Now().Format("20060102")
	expectedFilename := "daily-" + today + ".md"

	if !strings.HasSuffix(output, expectedFilename) {
		t.Errorf("Expected output to end with %s, got %s", expectedFilename, output)
	}

	// Verify file content
	content, err := os.ReadFile(output)
	if err != nil {
		t.Fatalf("Failed to read created file: %v", err)
	}

	expectedTitle := "Daily Note: " + time.Now().Format("2006-01-02")
	if !strings.Contains(string(content), "# "+expectedTitle) {
		t.Errorf("Content does not contain expected daily title. Got: %s", string(content))
	}

	if !strings.Contains(string(content), "tags: [daily]") {
		t.Errorf("Content does not contain expected daily tag. Got: %s", string(content))
	}
}
