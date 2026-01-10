# Plan: Daily Note Functionality & Command Rename

## Phase 1: Rename 'new' Command to 'create'

- [x] Task: Create reproduction test case for existing 'new' command behavior c52437c
    - *Goal:* Ensure we have a baseline test before refactoring.
    - *Details:* Verify 'new' creates a file with expected content.
- [ ] Task: Rename 'new.go' to 'create.go' and refactor command definition
    - *Goal:* Change the Cobra command Use string from "new" to "create".
    - *Details:* Update all internal references and help text.
- [ ] Task: Verify 'create' command passes baseline tests
    - *Goal:* Ensure the rename didn't break basic functionality.
    - *Details:* Update the test to call 'create' instead of 'new'.
- [ ] Task: Conductor - User Manual Verification 'Rename 'new' Command to 'create'' (Protocol in workflow.md)

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
