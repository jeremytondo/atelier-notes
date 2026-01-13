# Plan: Open New Note After Creation

## Phase 1: Configuration
- [ ] Task: Write tests for new config option in `config.lua`.
- [ ] Task: Update `lua/atelier-notes/config.lua` to include `ui.open_after_create = true`.
- [ ] Task: Conductor - User Manual Verification 'Phase 1: Configuration' (Protocol in workflow.md)

## Phase 2: Core Logic Enhancement
- [ ] Task: Write tests for `core.create_note` to ensure it returns the trimmed file path.
- [ ] Task: Update `lua/atelier-notes/core.lua` to trim whitespace from the CLI output.
- [ ] Task: Conductor - User Manual Verification 'Phase 2: Core Logic Enhancement' (Protocol in workflow.md)

## Phase 3: UI Integration (Open/Notify)
- [ ] Task: Write tests for `init.create_note` to verify "edit" or "notify" behavior based on config.
- [ ] Task: Update `lua/atelier-notes/init.lua` to handle opening or notifying based on `ui.open_after_create`.
- [ ] Task: Conductor - User Manual Verification 'Phase 3: UI Integration (Open/Notify)' (Protocol in workflow.md)
