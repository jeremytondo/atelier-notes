# Atelier Notes

Set of tools for working with local markdown notes.

## Usage

### Create a new note
To create a new note with a specific title:
```bash
atelier-notes create "My Note Title"
```
This will create a file named `my-note-title.md` in your configured notes directory.

### Create a daily note
To create a note for today:
```bash
atelier-notes create --daily
```
This will create a file named `daily-YYYYMMDD.md` with the `#daily` tag and a standardized title.

### Ask about your notes
To ask a natural language question about your notes:
```bash
atelier-notes ask "What are the main projects I'm working on?"
```
This command gathers context from all your markdown notes and uses AI to generate a concise answer.

### Flags
*   `-d, --dir`: (Global) Specify the directory to create/list notes in (overrides configuration).
*   `--daily`: (Create only) Create a daily note instead of a custom one.

## Neovim Plugin

This repository also contains a Neovim plugin that wraps the CLI for a seamless experience.

### Installation

Use your preferred package manager (e.g., `lazy.nvim`). Point it to your local directory or the repository URL.

**Using `lazy.nvim` (Local Development):**

```lua
{
  "jeremytondo/atelier-notes",
  dir = "~/path/to/atelier-notes", -- Adjust path as needed
  config = function()
    require("atelier-notes").setup({})
  end,
  -- Optional: Build the binary automatically
  build = "go build -o atelier-notes ./cmd/atelier-notes"
}
```

### Features

*   **:AtelierNotesCreate**: Create a new note with a title prompt.
