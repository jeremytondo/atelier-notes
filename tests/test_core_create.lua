-- tests/test_core_create.lua
vim.opt.rtp:prepend(vim.fn.getcwd())
local core = require("atelier-notes.core")

-- Mock vim.system
local original_system = vim.system
local mock_stdout = "/path/to/note.md\n"
local captured_args = nil

vim.system = function(args, opts)
    captured_args = args
    return {
        wait = function()
            return {
                code = 0,
                stdout = mock_stdout,
                stderr = ""
            }
        end
    }
end

-- Test create_note
local result = core.create_note("Test Note", { binary_path = "atelier-notes" })

-- Check if whitespace is trimmed
if result ~= "/path/to/note.md" then
    print(string.format("Failure: Expected '/path/to/note.md', got '%s'", result))
    os.exit(1)
end

-- Check if correct command was called (fix for 'new' vs 'create')
if captured_args[2] ~= "create" then
    print(string.format("Failure: Expected command 'create', got '%s'", captured_args[2]))
    os.exit(1)
end

print("Success: Core logic is correct")
os.exit(0)

