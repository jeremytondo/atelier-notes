package notes

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

// CreateNote creates a new markdown note with the given title in the target directory.
// It returns the absolute path of the created file.
func CreateNote(title string, targetDir string) (string, error) {
	slug := generateSlug(title)
	filename := fmt.Sprintf("%s.md", slug)
	return createNote(title, filename, targetDir, []string{})
}

// CreateDailyNote creates a new daily note in the target directory.
func CreateDailyNote(targetDir string) (string, error) {
	now := time.Now()
	filename := fmt.Sprintf("daily-%s.md", now.Format("20060102"))
	title := fmt.Sprintf("Daily Note: %s", now.Format("2006-01-02"))
	return createNote(title, filename, targetDir, []string{"daily"})
}

func createNote(title, filename, targetDir string, tags []string) (string, error) {
	// Ensure target directory exists
	if err := os.MkdirAll(targetDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create directory: %w", err)
	}

	fullPath := filepath.Join(targetDir, filename)
	absPath, err := filepath.Abs(fullPath)
	if err != nil {
		return "", fmt.Errorf("failed to resolve absolute path: %w", err)
	}

	// Check if file already exists to prevent overwrite
	if _, err := os.Stat(absPath); err == nil {
		return "", fmt.Errorf("file already exists: %s", absPath)
	}

	// Generate Content (Frontmatter)
	content := generateFrontmatter(title, tags)

	// Write File
	if err := os.WriteFile(absPath, []byte(content), 0644); err != nil {
		return "", fmt.Errorf("failed to write file: %w", err)
	}

	return absPath, nil
}

func generateSlug(title string) string {
	// Lowercase
	s := strings.ToLower(title)
	// Remove invalid chars (keep alphanumerics and spaces)
	reg := regexp.MustCompile(`[^a-z0-9\s-]`)
	s = reg.ReplaceAllString(s, "")
	// Replace spaces with dashes
	s = strings.ReplaceAll(s, " ", "-")
	// Remove duplicate dashes
	regDash := regexp.MustCompile(`-+`)
	s = regDash.ReplaceAllString(s, "-")
	// Trim dashes
	s = strings.Trim(s, "-")
	return s
}

func generateFrontmatter(title string, tags []string) string {
	now := time.Now().Format("2006-01-02")
	tagsStr := "[]"
	if len(tags) > 0 {
		tagsStr = "[" + strings.Join(tags, ", ") + "]"
	}
	return fmt.Sprintf(`---
date: %s
tags: %s
---

# %s

`, now, tagsStr, title)
}
