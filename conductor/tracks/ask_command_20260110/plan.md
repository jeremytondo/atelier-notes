# Plan: AI "Ask" Command Implementation

## Phase 1: Foundation & Context Gathering [checkpoint: adc96b4]
- [x] Task: Create the `ask` command skeleton in `internal/cli/ask.go` and register it in `cli.go`.
- [x] Task: Implement `internal/notes/reader.go` to gather and concatenate all `.md` files from the notes directory.
- [x] Task: Implement unit tests for note aggregation, ensuring it handles empty directories and non-markdown files correctly.
- [x] Task: Conductor - User Manual Verification 'Phase 1: Foundation & Context Gathering' (Protocol in workflow.md)

## Phase 2: Gemini CLI Integration [checkpoint: e5d74ae]
- [x] Task: Implement the prompt construction logic (System Preamble + Notes Context + User Question). 3197190
- [x] Task: Implement a runner for the Gemini CLI to execute "headless" commands and capture output. 853d17c
- [x] Task: Integrate the prompt builder with the Gemini runner inside the `ask` command. 617eeb9
- [x] Task: Add basic stdout output for the AI's response. 617eeb9
- [x] Task: Conductor - User Manual Verification 'Phase 2: Gemini CLI Integration' (Protocol in workflow.md)

## Phase 3: Error Handling & UX
- [x] Task: Implement error handling to intercept Gemini CLI failures and provide user-friendly messages (especially for token limit/context window errors). d7c3c4f
- [~] Task: Add a "Thinking..." or "Processing..." status indicator to provide user feedback during the AI call.
- [ ] Task: Perform a final quality gate check (linting, basic manual verification, ensuring no tests were skipped).
- [ ] Task: Conductor - User Manual Verification 'Phase 3: Error Handling & UX' (Protocol in workflow.md)
