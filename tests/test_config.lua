-- tests/test_config.lua
vim.opt.rtp:prepend(vim.fn.getcwd())
local config = require("atelier-notes.config")

if config.defaults.ui.open_after_create ~= true then
    print("Failure: ui.open_after_create should default to true")
    os.exit(1)
end

print("Success: Config defaults are correct")
os.exit(0)
