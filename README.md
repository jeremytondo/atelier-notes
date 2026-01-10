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

### Flags
*   `-d, --dir`: (Global) Specify the directory to create/list notes in (overrides configuration).
*   `--daily`: (Create only) Create a daily note instead of a custom one.