package notes

import (
	"os"
	"path/filepath"
	"testing"
)

func TestListNotes(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "list-notes-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create a few test notes
	note1 := `---
date: 2026-01-01
tags: [tag1]
---

# Note One
`
	note2 := `# Note Two Without Frontmatter`

	if err := os.WriteFile(filepath.Join(tempDir, "note1.md"), []byte(note1), 0644); err != nil {
		t.Fatalf("Failed to write note1: %v", err)
	}
	if err := os.WriteFile(filepath.Join(tempDir, "note2.md"), []byte(note2), 0644); err != nil {
		t.Fatalf("Failed to write note2: %v", err)
	}
	// Add a non-markdown file to ensure it's skipped
	if err := os.WriteFile(filepath.Join(tempDir, "random.txt"), []byte("not a note"), 0644); err != nil {
		t.Fatalf("Failed to write random.txt: %v", err)
	}

	notes, err := ListNotes(tempDir)
	if err != nil {
		t.Fatalf("ListNotes failed: %v", err)
	}

	if len(notes) != 2 {
		t.Errorf("Expected 2 notes, got %d", len(notes))
	}

	// Verify details
	foundOne := false
	foundTwo := false
	for _, n := range notes {
		t.Logf("Found note: title=%q, path=%q", n.Title, n.Path)
		if n.Title == "Note One" {
			foundOne = true
			if n.Date != "2026-01-01" {
				t.Errorf("Note One has wrong date: %s", n.Date)
			}
			if len(n.Tags) != 1 || n.Tags[0] != "tag1" {
				t.Errorf("Note One has wrong tags: %v", n.Tags)
			}
		}
		if n.Title == "Note Two Without Frontmatter" {
			foundTwo = true
		}
	}

	if !foundOne {
		t.Error("Did not find Note One")
	}
	if !foundTwo {
		t.Error("Did not find Note Two")
	}
}

func TestParseNote_Fallback(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "parse-note-fallback")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	path := filepath.Join(tempDir, "no-title.md")
	if err := os.WriteFile(path, []byte("just some content"), 0644); err != nil {
		t.Fatalf("Failed to write file: %v", err)
	}

	note, err := parseNote(path)
	if err != nil {
		t.Fatalf("parseNote failed: %v", err)
	}

	if note.Title != "no-title" {
		t.Errorf("Expected title 'no-title' (filename fallback), got %s", note.Title)
	}
}
