package cli

import (
	"github.com/spf13/cobra"
)

var askCmd = &cobra.Command{
	Use:   "ask <question>",
	Short: "Ask a question about your notes",
	Long:  `Ask a natural language question about the contents of your notes. The command will gather context and use the Gemini CLI to generate an answer.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("Ask command is not yet implemented")
	},
}

func init() {
	rootCmd.AddCommand(askCmd)
}
