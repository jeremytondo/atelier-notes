# Spec: Open New Note After Creation

## Overview
This track enhances the note creation workflow by automatically opening the newly created note in the current Neovim window. This provides a more seamless transition from capture to editing.

## Functional Requirements
- **Automatic Opening:** Upon successful execution of `atelier-notes new "<title>"`, the plugin should read the file path returned on `stdout`.
- **Buffer Management:** The plugin must open the file at the returned path in the current window using standard Neovim API calls (e.g., `vim.cmd.edit`).
- **Configuration:** 
    - Add an option `ui.open_after_create` (boolean, default: `true`) to the plugin configuration.
    - If `true`, the note is opened immediately.
    - If `false`, the plugin displays a success notification confirming the note creation and showing the file path.

## Non-Functional Requirements
- **Reliability:** If the CLI returns a malformed path or the file cannot be opened, the plugin should notify the user of the error but not crash.
- **Performance:** Opening the note should be instantaneous once the CLI command completes.

## Acceptance Criteria
- Executing `:AtelierNotesCreate` and entering a title results in the new note being opened in the current window (default behavior).
- If `ui.open_after_create` is set to `false`, the note is created but NOT opened; instead, a notification appears: "Note created: <path>".
- Error notifications are shown if the file path is invalid or the CLI fails.

## Out of Scope
- Opening notes in splits or tabs (fixed to current window for this track).
- Opening existing notes (this is a creation-flow enhancement only).
