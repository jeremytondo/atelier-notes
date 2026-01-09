package cli

import (
	"fmt"
	"os"

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
		
		// Default to current directory if not specified or empty
		if targetDir == "" {
			var err error
			targetDir, err = os.Getwd()
			if err != nil {
				fmt.Printf("Error getting current directory: %v\n", err)
				os.Exit(1)
			}
		}

		path, err := notes.CreateNote(title, targetDir)
		if err != nil {
			fmt.Printf("Error creating note: %v\n", err)
			os.Exit(1)
		}

		// Print only the path to stdout for easy pipe/editor integration
		fmt.Println(path)
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
	newCmd.Flags().StringVarP(&targetDir, "dir", "d", "", "Directory to create the note in")
}
