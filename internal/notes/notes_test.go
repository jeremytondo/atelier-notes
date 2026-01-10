package notes

import (
	"os"
	"strings"
	"testing"
)

func TestCreateNote(t *testing.T) {
	// Setup temporary directory
	tempDir, err := os.MkdirTemp("", "notes-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	title := "Test Note"
	expectedSlug := "test-note"
	expectedFilename := expectedSlug + ".md"

	absPath, err := CreateNote(title, tempDir)
	if err != nil {
		t.Fatalf("CreateNote failed: %v", err)
	}

	// Verify path
	if !strings.HasSuffix(absPath, expectedFilename) {
		t.Errorf("Expected path to end with %s, got %s", expectedFilename, absPath)
	}

	// Verify file existence
	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		t.Errorf("File was not created at %s", absPath)
	}

	// Verify content
	content, err := os.ReadFile(absPath)
	if err != nil {
		t.Fatalf("Failed to read created file: %v", err)
	}

	if !strings.Contains(string(content), "# Test Note") {
		t.Errorf("Content does not contain expected title header. Got: %s", string(content))
	}

	// Verify file exists check
	_, err = CreateNote(title, tempDir)
	if err == nil {
		t.Error("Expected error when creating existing note, got nil")
	} else if !strings.Contains(err.Error(), "file already exists") {
		t.Errorf("Expected 'file already exists' error, got: %v", err)
	}
}

func TestGenerateSlug(t *testing.T) {
	tests := []struct {
		title    string
		expected string
	}{
		{"Hello World", "hello-world"},
		{"My-Awesome Note!", "my-awesome-note"},
		{"  Spaces  Everywhere  ", "spaces-everywhere"},
		{"SpecialChars@#$%^&*()", "specialchars"},
		{"Multiple---Dashes", "multiple-dashes"},
	}

	for _, tt := range tests {
		result := generateSlug(tt.title)
		if result != tt.expected {
			t.Errorf("generateSlug(%q) = %q; want %q", tt.title, result, tt.expected)
		}
	}
}
