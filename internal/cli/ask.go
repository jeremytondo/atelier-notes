package cli

import (
	"strings"

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

		cmd.PrintErrln("Thinking...")
		response, err := gemini.Run(prompt)
		if err != nil {
			errStr := err.Error()
			if strings.Contains(errStr, "token limit exceeded") || strings.Contains(errStr, "context window") {
				cmd.PrintErrln("The combined size of your notes exceeds the AI's current context window.")
				cmd.PrintErrln("Try removing some large notes or wait for a future update with smarter retrieval.")
			} else {
				cmd.PrintErrf("Gemini CLI failed: %v\n", err)
			}
			return
		}

		cmd.Println(response)
	},
}

func init() {
	rootCmd.AddCommand(askCmd)
}
