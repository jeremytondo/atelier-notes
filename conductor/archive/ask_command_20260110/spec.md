# Specification - AI "Ask" Command

## Overview
Implement a new `ask` command for the `atelier-notes` CLI. This command allows users to perform natural language queries against their entire local knowledge base. It uses the Gemini CLI in a "headless" mode to process a combined context of all notes, a system-defined persona, and the user's specific question.

## Functional Requirements
*   **Subcommand:** Implement `atelier-notes ask "<question>"` using Cobra.
*   **Context Gathering:**
    *   Iterate through the configured notes directory.
    *   Read the contents of all `.md` files.
*   **Prompt Construction:**
    *   **System Preamble:** Prepend a default instruction set (persona) to the AI. Example: "You are the Atelier Assistant, an expert at organizing and summarizing the user's notes. Use the following context to answer the question accurately and concisely."
    *   **Note Context:** Append the full content of all discovered notes.
    *   **User Question:** Append the specific question provided by the user.
*   **AI Integration:**
    *   Invoke the Gemini CLI via a headless shell command.
    *   Pass the constructed prompt as input.
*   **Output:**
    *   Wait for the complete response from Gemini CLI.
    *   Print the response directly to `stdout`.
*   **Error Handling:**
    *   If the Gemini CLI returns an error (e.g., token limit exceeded, network error), catch it and display a user-friendly message rather than a raw stack trace.
    *   Specifically handle context window overflows with a message suggesting the user prunes notes or waits for future retrieval-based updates.

## Non-Functional Requirements
*   **Simplicity:** Adhere to the "Path A" approach for all initial design decisions to minimize architectural complexity.
*   **Performance:** Provide feedback (e.g., a simple "Thinking..." message) while waiting for the AI response to avoid the appearance of a hung process.

## Acceptance Criteria
*   The command `atelier-notes ask "What are my main projects right now?"` successfully returns a relevant answer based on note content.
*   The command handles a lack of notes gracefully (e.g., "No notes found to search").
*   Errors from the Gemini CLI are intercepted and presented clearly to the user.

## Out of Scope
*   Semantic search or vector indexing (to be handled in a future track).
*   Streaming responses.
*   Advanced persona customization via config or skills.
*   Selective note filtering (e.g., tags or specific files).
