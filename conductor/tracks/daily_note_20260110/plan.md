# Plan: Daily Note Functionality & Command Rename

## Phase 1: Rename 'new' Command to 'create' [checkpoint: 5246e6d]

- [x] Task: Create reproduction test case for existing 'new' command behavior c52437c
- [x] Task: Rename 'new.go' to 'create.go' and refactor command definition 5ad2b6d
- [x] Task: Verify 'create' command passes baseline tests 5ad2b6d
- [x] Task: Conductor - User Manual Verification 'Rename 'new' Command to 'create'' (Protocol in workflow.md) 5246e6d

## Phase 2: Implement Daily Note Logic

- [ ] Task: Write failing test for '--daily' flag
    - *Goal:* Define the expected behavior for 'create --daily'.
    - *Details:* Test should assert:
        - Filename matches 'daily-YYYYMMDD.md'
        - Content contains 'Daily Note: YYYY-MM-DD'
        - Content contains '#daily' tag
- [ ] Task: Implement '--daily' flag in 'create' command
    - *Goal:* Add the boolean flag to the Cobra command.
- [ ] Task: Implement Daily Note generation logic
    - *Goal:* Pass the failing test.
    - *Details:*
        - Logic to generate current date strings.
        - Logic to construct filename and content based on flag.
        - Ensure it handles the case where no title argument is provided if --daily is set.
- [ ] Task: Conductor - User Manual Verification 'Implement Daily Note Logic' (Protocol in workflow.md)

## Phase 3: Final Polish & Documentation

- [ ] Task: Update README to reflect 'create' command and '--daily' usage
- [ ] Task: Run full test suite and ensure >80% coverage
- [ ] Task: Conductor - User Manual Verification 'Final Polish & Documentation' (Protocol in workflow.md)
