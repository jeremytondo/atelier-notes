package cli

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/jeremytondo/atelier-notes/internal/gemini"
)

func TestAskCommand(t *testing.T) {
	// Setup temp dir for notes
	tempDir, err := os.MkdirTemp("", "cli-ask-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer func() { _ = os.RemoveAll(tempDir) }()

	// Create a test note
	notePath := filepath.Join(tempDir, "test.md")
	err = os.WriteFile(notePath, []byte("# Test Note\nThis is a test."), 0644)
	if err != nil {
		t.Fatalf("Failed to create test note: %v", err)
	}

	// Set targetDir to tempDir
	targetDir = tempDir

	// Mock gemini
	gemini.ExecCommand = func(command string, args ...string) *exec.Cmd {
		cs := []string{"-test.run=TestHelperProcess", "--", command}
		cs = append(cs, args...)
		cmd := exec.Command(os.Args[0], cs...)
		cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1"}
		return cmd
	}
	defer func() { gemini.ExecCommand = exec.Command }()

	// Buffer to capture stdout
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	// Execute command: atelier-notes ask "What is in the test note?"
	rootCmd.SetArgs([]string{"ask", "What is in the test note?"})
	err = rootCmd.Execute()
	if err != nil {
		t.Fatalf("Command execution failed: %v", err)
	}

	output := strings.TrimSpace(buf.String())
	expected := "mocked output"
	if !strings.Contains(output, expected) {
		t.Errorf("Expected output to contain %q, got %q", expected, output)
	}
	if !strings.Contains(output, "Thinking...") {
		t.Errorf("Expected output to contain 'Thinking...', got %q", output)
	}
}

func TestAskCommand_Error(t *testing.T) {
	// Setup temp dir for notes
	tempDir, err := os.MkdirTemp("", "cli-ask-error-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer func() { _ = os.RemoveAll(tempDir) }()

	// Set targetDir to tempDir
	targetDir = tempDir

	// Mock gemini to return an error
	gemini.ExecCommand = func(command string, args ...string) *exec.Cmd {
		cs := []string{"-test.run=TestHelperErrorProcess", "--", command}
		cs = append(cs, args...)
		cmd := exec.Command(os.Args[0], cs...)
		cmd.Env = []string{"GO_WANT_HELPER_ERROR_PROCESS=1"}
		return cmd
	}
	defer func() { gemini.ExecCommand = exec.Command }()

	// Buffer to capture stdout/stderr
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	// Execute command
	rootCmd.SetArgs([]string{"ask", "some question"})
	err = rootCmd.Execute()
	// The command should not return an error because we handle it and print a friendly message
	if err != nil {
		t.Fatalf("Command execution failed: %v", err)
	}

	output := buf.String()
	if !strings.Contains(output, "exceeds the AI's current context window") {
		t.Errorf("Expected context window error message, got %q", output)
	}
}

func TestHelperErrorProcess(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_ERROR_PROCESS") != "1" {
		return
	}
	fmt.Fprintf(os.Stderr, "token limit exceeded")
	os.Exit(1)
}

func TestHelperProcess(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}
	fmt.Print("mocked output")
	os.Exit(0)
}

func TestAskCommand_NoArgs(t *testing.T) {
	// Buffer to capture stdout/stderr
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	// Execute command: atelier-notes ask
	// Should fail because it requires an argument
	rootCmd.SetArgs([]string{"ask"})
	err := rootCmd.Execute()
	
	if err == nil {
		t.Error("Expected error when no question provided, got nil")
	}
}
