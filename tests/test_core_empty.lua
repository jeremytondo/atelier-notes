-- tests/test_core_empty.lua
vim.opt.rtp:prepend(vim.fn.getcwd())
local core = require("atelier-notes.core")

-- Mock vim.system to return empty stdout (or whitespace)
vim.system = function(args, opts)
    return {
        wait = function()
            return {
                code = 0,
                stdout = "   \n  ", -- whitespace only
                stderr = ""
            }
        end
    }
end

-- Test create_note
local result = core.create_note("Test Note", { binary_path = "atelier-notes" })

if result ~= nil then
    print(string.format("Failure: Expected nil for empty output, got '%s'", tostring(result)))
    os.exit(1)
end

print("Success: Core logic handles empty output correctly")
os.exit(0)

