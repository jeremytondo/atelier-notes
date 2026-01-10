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