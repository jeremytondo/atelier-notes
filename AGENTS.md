# Atelier Notes - Project Context

## Rules
- Keep it simple. Code should be easy to understand and read.
- Follow Go best practices
- Don't ever use numbered steps in comments. It's unnecessary and quickly gets our of date.
- Don't write tests. We don't need then.

## Project Overview
Atelier Notes is a CLI tool designed to manage a local-first, AI-enhanced knowledge base. It facilitates note creation, organization, and retrieval using a flat file structure (Markdown) and semantic search capabilities.

## Technical Architecture
- **Language:** Go (Golang)
- **CLI Framework:** `spf13/cobra`
- **Storage:** Local filesystem, flat directory of Markdown files.
- **Key Directories:**
  - `cmd/atelier-notes`: Entry point.
  - `internal/`: Core logic (CLI handlers, note processing).
  - `.blueprints/`: Architectural documentation.

## Core Workflows
- **Capture:** Quick capture via Google Keep (mobile) or Neovim (desktop).
- **Organization:** AI agents manage metadata and file structure.
- **Retrieval:** Semantic search via the Gemini CLI.

## Coding Conventions
- Follow standard Go project layout.
- Use `internal` package for non-exported logic.
- Error handling: Explicit checks, avoid panic in libraries.
