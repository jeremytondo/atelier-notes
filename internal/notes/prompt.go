package notes

import "fmt"

const DefaultSystemPreamble = `You are the Atelier Assistant, an expert at organizing and summarizing the user's notes. Use the following context to answer the question accurately and concisely.`

// BuildAskPrompt constructs the full prompt to be sent to the AI.
func BuildAskPrompt(notesContext, question string) string {
	return fmt.Sprintf("%s\n\nContext:\n%s\n\nQuestion: %s", DefaultSystemPreamble, notesContext, question)
}
