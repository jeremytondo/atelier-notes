package cli

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/jeremytondo/atelier-notes/internal/notes"
)

func TestListCommand(t *testing.T) {
	// Setup temp dir for notes
	tempDir, err := os.MkdirTemp("", "cli-list-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create a test note using internal logic to be sure it's valid
	_, err = notes.CreateNote("List Test Note", tempDir)
	if err != nil {
		t.Fatalf("Failed to create test note: %v", err)
	}

	// Buffer to capture stdout
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(new(bytes.Buffer))

	// Execute command: atelier-notes list --dir <tempDir>
	rootCmd.SetArgs([]string{"list", "--dir", tempDir})
	err = rootCmd.Execute()
	if err != nil {
		t.Fatalf("Command execution failed: %v", err)
	}

	// Parse JSON output
	var noteList []notes.Note
	if err := json.Unmarshal(buf.Bytes(), &noteList); err != nil {
		t.Fatalf("Failed to parse JSON output: %v. Output was: %s", err, buf.String())
	}

	if len(noteList) != 1 {
		t.Errorf("Expected 1 note in list, got %d", len(noteList))
	}

	if noteList[0].Title != "List Test Note" {
		t.Errorf("Expected title 'List Test Note', got %s", noteList[0].Title)
	}

	if !filepath.IsAbs(noteList[0].Path) {
		t.Errorf("Expected absolute path, got %s", noteList[0].Path)
	}
}
