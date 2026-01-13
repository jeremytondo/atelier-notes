-- tests/test_init_create.lua
vim.opt.rtp:prepend(vim.fn.getcwd())

-- Mock dependencies
package.loaded["atelier-notes.core"] = {}
local core = require("atelier-notes.core")

-- Setup mocks
local edit_called_with = nil
local notify_called_with = nil

vim.cmd = {
    edit = function(path)
        edit_called_with = path
    end
}

vim.notify = function(msg)
    notify_called_with = msg
end

-- Mock input
vim.ui.input = function(opts, on_confirm)
    on_confirm("Test Note")
end

-- Load module
local atelier = require("atelier-notes.init")

-- Test Case 1: open_after_create = true (default)
core.create_note = function(title, opts)
    return "/path/to/note1.md"
end

edit_called_with = nil
notify_called_with = nil

-- Ensure default config
atelier.setup({})
atelier.create_note()

if edit_called_with ~= "/path/to/note1.md" then
    print("Failure (Case 1): Expected vim.cmd.edit('/path/to/note1.md')")
    os.exit(1)
end

-- Test Case 2: open_after_create = false
edit_called_with = nil
notify_called_with = nil

core.create_note = function(title, opts)
    return "/path/to/note2.md"
end

atelier.setup({
    ui = {
        open_after_create = false
    }
})
atelier.create_note()

if edit_called_with ~= nil then
    print("Failure (Case 2): Should not open file when open_after_create = false")
    os.exit(1)
end

if not notify_called_with or not string.find(notify_called_with, "/path/to/note2.md") then
    print("Failure (Case 2): Expected notification with path")
    os.exit(1)
end

print("Success: UI integration logic is correct")
os.exit(0)
