package notes

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestReadAllNotes(t *testing.T) {
	// Setup temp dir
	tempDir, err := os.MkdirTemp("", "notes-reader-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create some test files
	files := map[string]string{
		"note1.md":      "# Note 1\nContent 1",
		"note2.md":      "# Note 2\nContent 2",
		"ignored.txt":   "Should be ignored",
		"subdir/sub.md": "# Sub Note\nContent Sub", // Should handle subdirs if we decide to recurse, but spec said "flat file structure" in product.md.
		// Wait, product.md says "flat directory of Markdown files". But let's check if we want recursion.
		// Spec says "Iterate through the configured notes directory". Usually implies flat or recursive.
		// For now, I'll assume flat based on "flat file structure", but standard behavior usually ignores subdirs if flat.
		// Let's stick to flat for now as per product.md.
	}

	for name, content := range files {
		path := filepath.Join(tempDir, name)
		if strings.Contains(name, "/") {
			err := os.MkdirAll(filepath.Dir(path), 0755)
			if err != nil {
				t.Fatalf("Failed to create subdir: %v", err)
			}
		}
		err := os.WriteFile(path, []byte(content), 0644)
		if err != nil {
			t.Fatalf("Failed to write file %s: %v", name, err)
		}
	}

	// Test case 1: read all markdown files
	content, err := ReadAllNotes(tempDir)
	if err != nil {
		t.Fatalf("ReadAllNotes failed: %v", err)
	}

	// Check if content contains expected strings
	expected := []string{"# Note 1", "Content 1", "# Note 2", "Content 2"}
	for _, exp := range expected {
		if !strings.Contains(content, exp) {
			t.Errorf("Expected content to contain '%s'", exp)
		}
	}

	// Check if it ignored non-md files
	if strings.Contains(content, "Should be ignored") {
		t.Error("Content should not contain text files")
	}

	// Check format/delimiters (optional but good for context)
	// We might want to prepend filename or something.
}

func TestReadAllNotes_Empty(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "notes-reader-empty")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	content, err := ReadAllNotes(tempDir)
	if err != nil {
		t.Fatalf("ReadAllNotes failed: %v", err)
	}

	if content != "" {
		t.Errorf("Expected empty content, got '%s'", content)
	}
}
