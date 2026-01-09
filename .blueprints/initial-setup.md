# Initial Setup Plan: Atelier Notes CLI

This document outlines the initial scaffolding and development plan for the `atelier-notes` Go CLI tool.

## 1. Project Initialization

- **Module Name:** `github.com/jeremytondo/atelier-notes`
- **Core Library:** `github.com/spf13/cobra` for CLI command management.

## 2. Directory Structure

Following standard Go conventions (`cmd`/`internal`) and specific package naming preferences:

```text
atelier-notes/
├── go.mod                # Module definition
├── go.sum                # Dependency checksums
├── cmd/
│   └── atelier-notes/
│       └── main.go       # Entry point
└── internal/
    ├── cli/
    │   ├── cli.go        # CLI package entry (Root command)
    │   └── new.go        # 'new' command definition
    └── notes/
        └── notes.go      # Core note management logic
```

## 3. Implementation Phases

### Phase 1: Scaffolding
- Initialize the Go module.
- Create the directory hierarchy.
- Set up the base Cobra `root` command in `internal/cli/cli.go`.

### Phase 2: Core Logic (`internal/notes/notes.go`)
- Implement `CreateNote` function.
- **Slug Generation:** Convert titles to kebab-case (e.g., "My Note" -> "my-note.md").
- **Frontmatter:** Generate YAML metadata (date, tags, status).
- **File I/O:** Handle file creation in the target directory.

### Phase 3: CLI Integration (`internal/cli/new.go`)
- Wire the `new` command to the `CreateNote` logic.
- Add flags for `--dir` (target directory) and `--tags`.
- Ensure the command outputs the absolute path of the created file for editor integration.

## 4. Immediate Goal: The `new` Command

Expected usage:
```bash
atelier-notes new "Project Gemini Planning" --dir ./Notes
```
- **Output:** Creates `./Notes/project-gemini-planning.md` (or similar slug format).
- **Success Criteria:** A markdown file is created with correct YAML frontmatter and the CLI returns the file path.
