# Specification: Daily Note Functionality & Command Rename

## Context
The project currently uses a `new` command to create notes. We want to rename this to `create` for better semantics. Additionally, we want to introduce a specialized workflow for "Daily Notes" which are standardized by date.

## Requirements

### 1. Rename Command
- **Current:** `atelier-notes new <title>`
- **New:** `atelier-notes create <title>`
- **Behavior:** The underlying functionality of creating a note remains the same, but the command entry point and help text must be updated.

### 2. Daily Note Feature
- **Flag:** Support a `--daily` flag for the `create` command.
- **Filename Format:** `daily-YYYYMMDD.md` (e.g., `daily-20260110.md`).
- **Title Format:** `Daily Note: YYYY-MM-DD` (e.g., `Daily Note: 2026-01-10`).
- **Tags:** Automatically append the tag `daily` to the note content using the existing metadata format.
- **Argument Handling:** If `--daily` is passed, the `<title>` argument should be ignored in favor of the standardized daily format.
- **Existing Note**: We should check to see if a daily note already exists. If it does, we should let the user know that it already exists and do nothing.

## User Stories
- As a user, I want to run `atelier-notes create --daily` to instantly generate a note for today so I can start journaling without worrying about naming conventions.
- As a user, I want the command to be named `create` instead of `new` because it feels more natural to me.

## Implementation Details
- **CLI Framework:** Update `cobra` command definitions.
- **File Logic:** Ensure date generation uses the user's local time.
- **Refactoring:** Rename `internal/cli/new.go` to `internal/cli/create.go`.
