package gemini

import (
	"bytes"
	"fmt"
	"os/exec"
)

var ExecCommand = exec.Command

// Run executes the Gemini CLI with the given prompt and returns the output.
func Run(prompt string) (string, error) {
	cmd := ExecCommand("gemini")

	cmd.Stdin = bytes.NewBufferString(prompt)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("gemini error: %v, stderr: %s", err, stderr.String())
	}

	return out.String(), nil
}
