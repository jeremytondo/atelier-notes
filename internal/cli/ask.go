package cli

import (
	"github.com/jeremytondo/atelier-notes/internal/config"
	"github.com/jeremytondo/atelier-notes/internal/gemini"
	"github.com/jeremytondo/atelier-notes/internal/notes"
	"github.com/spf13/cobra"
)

var askCmd = &cobra.Command{
	Use:   "ask <question>",
	Short: "Ask a question about your notes",
	Long:  `Ask a natural language question about the contents of your notes. The command will gather context and use the Gemini CLI to generate an answer.`, 
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		question := args[0]

		if targetDir == "" {
			targetDir = config.GetNotesDir()
		}

		notesContext, err := notes.ReadAllNotes(targetDir)
		if err != nil {
			cmd.PrintErrf("Error reading notes: %v\n", err)
			return
		}

		prompt := notes.BuildAskPrompt(notesContext, question)

		response, err := gemini.Run(prompt)
		if err != nil {
			cmd.PrintErrf("Error from AI: %v\n", err)
			return
		}

		cmd.Println(response)
	},
}

func init() {
	rootCmd.AddCommand(askCmd)
}