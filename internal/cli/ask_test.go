package cli

import (
	"bytes"
	"strings"
	"testing"
)

func TestAskCommand_Skeleton(t *testing.T) {
	// Buffer to capture stdout/stderr
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	// Execute command: atelier-notes ask "Test question"
	rootCmd.SetArgs([]string{"ask", "Test question"})
	err := rootCmd.Execute()
	
	// Expect nil error for now (skeleton should just run)
	if err != nil {
		t.Fatalf("Command execution failed: %v", err)
	}

	// Verify output contains a placeholder message
	output := strings.TrimSpace(buf.String())
	expected := "Ask command is not yet implemented"
	if !strings.Contains(output, expected) {
		t.Errorf("Expected output to contain '%s', got '%s'", expected, output)
	}
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
