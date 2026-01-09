package notes

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

type Note struct {
	Title string   `json:"title"`
	Path  string   `json:"path"`
	Date  string   `json:"date"`
	Tags  []string `json:"tags"`
}

type frontmatter struct {
	Date string   `yaml:"date"`
	Tags []string `yaml:"tags"`
}

// ListNotes scans the given directory for markdown files and parses them.
func ListNotes(dir string) ([]Note, error) {
	var notes []Note

	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %w", err)
	}

	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".md" {
			continue
		}

		fullPath := filepath.Join(dir, entry.Name())
		note, err := parseNote(fullPath)
		if err != nil {
			// Skip files that can't be parsed or log warning? 
			// For now, we'll just skip to be robust.
			continue
		}
		notes = append(notes, note)
	}

	return notes, nil
}

func parseNote(path string) (Note, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return Note{}, err
	}

	note := Note{
		Path: path,
	}

	// Simple frontmatter parsing
	// Assumes file starts with ---
	parts := bytes.SplitN(content, []byte("---"), 3)
	if len(parts) >= 3 {
		// parts[0] is empty (before first ---)
		// parts[1] is frontmatter
		// parts[2] is content
		
		var fm frontmatter
		if err := yaml.Unmarshal(parts[1], &fm); err == nil {
			note.Date = fm.Date
			note.Tags = fm.Tags
		}
		
		// Parse title from content (first line starting with #)
		scanner := bufio.NewScanner(bytes.NewReader(parts[2]))
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			if strings.HasPrefix(line, "# ") {
				note.Title = strings.TrimPrefix(line, "# ")
				break
			}
		}
	} else {
        // Fallback if no frontmatter found, try to find title in whole file
        scanner := bufio.NewScanner(bytes.NewReader(content))
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			if strings.HasPrefix(line, "# ") {
				note.Title = strings.TrimPrefix(line, "# ")
				break
			}
		}
    }

	// Fallback title to filename if not found
	if note.Title == "" {
		filename := filepath.Base(path)
		note.Title = strings.TrimSuffix(filename, filepath.Ext(filename))
	}

	return note, nil
}
