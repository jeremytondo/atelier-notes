package gemini

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

// Mocking exec.Command as per common Go patterns
func TestRun_Mock(t *testing.T) {
	execCommand = func(command string, args ...string) *exec.Cmd {
		cs := []string{"-test.run=TestHelperProcess", "--", command}
		cs = append(cs, args...)
		cmd := exec.Command(os.Args[0], cs...)
		cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1"}
		return cmd
	}
	defer func() { execCommand = exec.Command }()

	out, err := Run("test prompt")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := "mocked output"
	if out != expected {
		t.Errorf("Expected %q, got %q", expected, out)
	}
}

func TestHelperProcess(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}
	fmt.Print("mocked output")
	os.Exit(0)
}