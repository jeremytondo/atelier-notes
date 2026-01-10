package notes

import (
	"strings"
	"testing"
)

func TestBuildAskPrompt(t *testing.T) {
	notesContext := "---\nNote: test.md ---\nContent of test note."
	question := "What is in the test note?"
	
prompt := BuildAskPrompt(notesContext, question)
	
	// Check for System Preamble
	if !strings.Contains(prompt, "You are the Atelier Assistant") {
		t.Error("Prompt missing system preamble")
	}
	
	// Check for Notes Context
	if !strings.Contains(prompt, notesContext) {
		t.Error("Prompt missing notes context")
	}
	
	// Check for User Question
	if !strings.Contains(prompt, question) {
		t.Error("Prompt missing user question")
	}
	
	// Check for structure
	if !strings.HasPrefix(prompt, "You are the Atelier Assistant") {
		t.Error("Prompt should start with system preamble")
	}
}
