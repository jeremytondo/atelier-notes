# Tech Stack - Atelier Notes

## Core Technologies
*   **Go (Golang) 1.23.0:** Chosen for its performance, strong typing, and excellent support for building CLI tools.
*   **Cobra:** The primary CLI framework, providing a robust structure for commands, subcommands, and flags.
*   **Viper:** Used for configuration management, with a preference for **YAML** for hand-edited configuration files.

## Libraries and Tools
*   **fsnotify:** For monitoring filesystem changes, potentially used for real-time indexing or synchronization.
*   **Afero:** A filesystem abstraction layer that enables easier testing by mocking the file system.
*   **Gemini CLI:** Integration for natural language querying (`ask` command), semantic search, and AI-driven features.

## Data Formats
*   **Markdown:** The primary format for user notes, ensuring readability and interoperability.
*   **YAML:** The standard format for user-facing configuration files.
*   **JSON:** The primary format for machine-readable output and automated integrations.

---

# Tech Stack - Atelier-Notes.nvim

## Core Technologies
- **Language:** Lua
- **Platform:** Neovim (Targeting the latest stable version)
- **Distribution Target:** LazyVim (Latest version)

## External Integrations
- **Core Binary:** `atelier-notes` (CLI tool)
    - The plugin acts as a wrapper and UI for this binary.

## Key Neovim Plugins & Libraries
- **Plugin Manager:** `lazy.nvim` (Assumed environment)
- **UI & Utilities:** `snacks.nvim`
    - **Pickers:** `snacks.picker` (Primary interface for selection and searching)
    - **General UI:** Used for inputs, notifications, and other widgets.