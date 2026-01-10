package gemini

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

// Mocking exec.Command as per common Go patterns
func TestRun_Mock(t *testing.T) {
	ExecCommand = func(command string, args ...string) *exec.Cmd {
		// Verify no args are passed (since we removed --headless)
		if len(args) > 0 {
			// This might be too strict if we add other flags later, but for now it ensures we removed the bad one.
			// Actually, let's just check that --headless is NOT in args.
			for _, arg := range args {
				if arg == "--headless" {
					t.Fatal("unexpected argument --headless")
				}
			}
		}

		cs := []string{"-test.run=TestHelperProcess", "--", command}
		cs = append(cs, args...)
		cmd := exec.Command(os.Args[0], cs...)
		cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1"}
		return cmd
	}
	defer func() { ExecCommand = exec.Command }()

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