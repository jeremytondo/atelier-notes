# Plan: Open New Note After Creation

## Phase 1: Configuration
- [x] Task: Write tests for new config option in `config.lua`.
- [x] Task: Update `lua/atelier-notes/config.lua` to include `ui.open_after_create = true`.
- [x] Task: Conductor - User Manual Verification 'Phase 1: Configuration' (Protocol in workflow.md)

## Phase 2: Core Logic Enhancement
- [x] Task: Write tests for `core.create_note` to ensure it returns the trimmed file path.
- [x] Task: Update `lua/atelier-notes/core.lua` to trim whitespace from the CLI output.
- [x] Task: Conductor - User Manual Verification 'Phase 2: Core Logic Enhancement' (Protocol in workflow.md)

## Phase 3: UI Integration (Open/Notify)
- [x] Task: Write tests for `init.create_note` to verify "edit" or "notify" behavior based on config.
- [x] Task: Update `lua/atelier-notes/init.lua` to handle opening or notifying based on `ui.open_after_create`.
- [x] Task: Conductor - User Manual Verification 'Phase 3: UI Integration (Open/Notify)' (Protocol in workflow.md)