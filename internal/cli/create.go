package cli

import (
	"os"

	"github.com/jeremytondo/atelier-notes/internal/config"
	"github.com/jeremytondo/atelier-notes/internal/notes"
	"github.com/spf13/cobra"
)

var (
	targetDir string
	daily     bool
)

var createCmd = &cobra.Command{
	Use:   "create [title]",
	Short: "Create a new note",
	Args: func(cmd *cobra.Command, args []string) error {
		if daily {
			return nil
		}
		return cobra.ExactArgs(1)(cmd, args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		// Default to configured directory if not specified
		if targetDir == "" {
			targetDir = config.GetNotesDir()
		}

		var path string
		var err error

		if daily {
			path, err = notes.CreateDailyNote(targetDir)
		} else {
			title := args[0]
			path, err = notes.CreateNote(title, targetDir)
		}

		if err != nil {
			cmd.PrintErrf("Error creating note: %v\n", err)
			os.Exit(1)
		}

		// Print only the path to stdout for easy pipe/editor integration
		cmd.Println(path)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().StringVarP(&targetDir, "dir", "d", "", "Directory to create the note in")
	createCmd.Flags().BoolVar(&daily, "daily", false, "Create a daily note")
}
