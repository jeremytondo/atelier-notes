package cli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/jeremytondo/atelier-notes/internal/config"
	"github.com/jeremytondo/atelier-notes/internal/notes"
	"github.com/spf13/cobra"
)

var useFzf bool

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all notes",
	Run: func(cmd *cobra.Command, args []string) {
		// 1. Get Directory
		dir := targetDir
		if dir == "" {
			dir = config.GetNotesDir()
		}

		// 2. Get Notes
		noteList, err := notes.ListNotes(dir)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error listing notes: %v\n", err)
			os.Exit(1)
		}

		// 3. Handle Output
		if useFzf {
			runFzf(cmd, noteList)
		} else {
			printJSON(cmd, noteList)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVar(&useFzf, "fzf", false, "Use fzf for fuzzy searching")
}

func printJSON(cmd *cobra.Command, notes []notes.Note) {
	encoder := json.NewEncoder(cmd.OutOrStdout())
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(notes); err != nil {
		cmd.PrintErrf("Error parsing JSON: %v\n", err)
		os.Exit(1)
	}
}

func runFzf(cmd *cobra.Command, noteList []notes.Note) {
	// check if fzf is installed
	if _, err := exec.LookPath("fzf"); err != nil {
		cmd.PrintErrln("Error: fzf is not installed or not in PATH")
		os.Exit(1)
	}

	// Prepare input for fzf
	// Format: Title <tab> Date <tab> Path
	var inputBuffer bytes.Buffer
	delimiter := "\t"
	
	for _, n := range noteList {
		if _, err := fmt.Fprintf(&inputBuffer, "%s%s%s%s%s\n", n.Title, delimiter, n.Date, delimiter, n.Path); err != nil {
			cmd.PrintErrf("Error writing to buffer: %v\n", err)
			return
		}
	}

	// Construct fzf command
	// --preview: attempt to use 'bat' (color) or 'cat'
	previewCmd := "cat {-1}"
	if _, err := exec.LookPath("bat"); err == nil {
		previewCmd = "bat --style=numbers --color=always {-1}"
	}

	fzfCmd := exec.Command("fzf", 
		"--delimiter", delimiter,
		"--with-nth", "1", // Show only Title
		"--preview", previewCmd,
		"--preview-window", "right:50%:wrap",
	)
	
	fzfCmd.Stdin = &inputBuffer
	fzfCmd.Stderr = os.Stderr // fzf uses stderr for UI

	output, err := fzfCmd.Output()
	if err != nil {
		// fzf returns non-zero if no match or cancelled (1 or 130)
		if exitErr, ok := err.(*exec.ExitError); ok {
			if exitErr.ExitCode() == 130 {
				// Cancelled by user
				return
			}
		}
		cmd.PrintErrf("Error running fzf: %v\n", err)
		return
	}

	// Parse selection to get path
	selectedLine := strings.TrimSpace(string(output))
	if selectedLine == "" {
		return
	}

	parts := strings.Split(selectedLine, delimiter)
	if len(parts) < 3 {
		cmd.PrintErrln("Error parsing selection")
		return
	}
	path := parts[len(parts)-1] // Last part is the path

	// Open editor
	editor := os.Getenv("EDITOR")
	if editor == "" {
		cmd.Println(path) // Fallback: just print path
		return
	}

	editorCmd := exec.Command(editor, path)
	editorCmd.Stdin = os.Stdin
	editorCmd.Stdout = os.Stdout
	editorCmd.Stderr = os.Stderr
	if err := editorCmd.Run(); err != nil {
		cmd.PrintErrf("Error opening editor: %v\n", err)
	}
}
