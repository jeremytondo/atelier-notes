package notes

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// ReadAllNotes reads all Markdown files in the given directory and returns their concatenated content.
func ReadAllNotes(dir string) (string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return "", fmt.Errorf("failed to read directory: %w", err)
	}

	var sb strings.Builder

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		if !strings.HasSuffix(strings.ToLower(entry.Name()), ".md") {
			continue
		}

		path := filepath.Join(dir, entry.Name())
		content, err := os.ReadFile(path)
		if err != nil {
			// Warn and skip, or return error?
			// For robustness, let's skip unreadable files but maybe log?
			// Since we don't have a logger setup, let's just skip.
			continue
		}

		sb.WriteString(fmt.Sprintf("\n--- Note: %s ---\n", entry.Name()))
		sb.Write(content)
		sb.WriteString("\n")
	}

	return sb.String(), nil
}
