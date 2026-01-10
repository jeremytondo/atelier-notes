package cli

import (
	"os"

	"github.com/jeremytondo/atelier-notes/internal/config"
	"github.com/jeremytondo/atelier-notes/internal/notes"
	"github.com/spf13/cobra"
)

var (
	targetDir string
)

var newCmd = &cobra.Command{
	Use:   "new [title]",
	Short: "Create a new note",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		title := args[0]
		
		// Default to configured directory if not specified
		if targetDir == "" {
			targetDir = config.GetNotesDir()
		}

		path, err := notes.CreateNote(title, targetDir)
		if err != nil {
			cmd.PrintErrf("Error creating note: %v\n", err)
			os.Exit(1)
		}

		// Print only the path to stdout for easy pipe/editor integration
		cmd.Println(path)
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
	newCmd.Flags().StringVarP(&targetDir, "dir", "d", "", "Directory to create the note in")
}
