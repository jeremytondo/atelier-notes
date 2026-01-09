# List Notes

We need a list command that allows listing available notes. I think there are two main ways we'd want to do this:

## Simple List
I think we want to be able to output a simple computer readable output that can be used by other tools. This should include details of each note. It should include the file path, the title, the tags, and the date. Would we maybe want to do this via a simple table? Maybe we'd want to use json?

## FZF
It would also be cool to use fzf to allow a user to fuzzy find notes and see a preview of them. I think fzf would be a good tool for doing this.

I think this command should like something like this:

This would get you the computer readable version.
```bash
atelier-notes list

This would get you a list opened up in fzf.
```
```bash
atlier-notes list --fzf
```

## FZF details
We should set up FZF to open in a nice readable way and allow it to show a preview of files that are using the default color scheme.

## Implementation Plan

1.  **Enhance `internal/notes` Package**:
    *   Add a `Note` struct to hold metadata (Path, Title, Date, Tags).
    *   Implement a `ListNotes(dir string)` function that:
        *   Scans the notes directory for `.md` files.
        *   Parses each file to extract the **Title** (from the first `# Header`), **Date**, and **Tags** (from YAML frontmatter).

2.  **Implement `list` Command (`internal/cli/list.go`)**:
    *   Create a new `listCmd`.
    *   **Default Output**: Output the list of notes in **JSON** format.
    *   **FZF Integration (`--fzf`)**:
        *   Check if `fzf` is installed.
        *   Format the list for `fzf` (e.g., display Title and Date).
        *   Configure `fzf`'s preview window to show the note content (using `bat` if available, otherwise `cat`).
        *   On selection, open the note in the default `$EDITOR` (or print the path if no editor is set).
