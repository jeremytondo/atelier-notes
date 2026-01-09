package cli

import (
	"os"

	"github.com/jeremytondo/atelier-notes/internal/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "atelier-notes",
	Short: "AI-enhanced note-taking system",
	Long: `Atelier Notes is a CLI tool for managing a local-first, 
AI-enhanced knowledge base. It handles note creation, organization, 
and retrieval.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(config.Load)
}
