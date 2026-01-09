# AI Notes Workflow

## General Overview

The core vision is to build a "local-first" yet cloud-accessible AI-enhanced note-taking system. The primary goal is to minimize friction during capture and retrieval while leveraging AI to handle organization and synthesis.

**Key Components:**
*   **The Brain:** Gemini CLI acts as the central intelligence for organizing, summarizing, and retrieving information.
*   **The Editor:** Neovim (LazyVim) serves as the primary tool for writing and deep editing.
*   **The Bridge:** Google Drive (via Desktop) provides a synchronized filesystem that allows local text manipulation while offering cloud accessibility.

The system aims to replicate the ease of Google Keep for capturing thoughts and the power of a Markdown knowledge base for long-term storage, glued together by AI agents that perform the roles of "Secretary" and "Librarian."

---

## Implementation Details

### 1. Capturing Notes

We employ a multi-modal approach to capture information depending on the context.

*   **Mobile / Quick Capture:**
    *   **Tool:** Google Keep.
    *   **Workflow:** Use Keep for voice-to-text, quick sticky notes, or checklist items on the go. These serve as temporary "jots" rather than permanent records.
    *   **Processing:** An AI agent will periodically scan these captures and migrate them to the permanent storage.

*   **Desktop / Neovim Workflow:**
    *   **Tool:** Neovim (LazyVim).
    *   **Workflow:** Direct creation of Markdown files using a custom shell script (e.g., `new-note.sh`) or keybindings.
    *   **Automation:**
        *   Script prompts for a "Title" (e.g., "Project Gemini Planning").
        *   Automatically converts title to a filename slug (e.g., `project-gemini-planning.md`).
        *   Pre-fills YAML frontmatter (date, tags, status).

### 2. Organizing Notes

Organization is flat, file-based, and maintained by AI agents to prevent manual overhead.

*   **Structure:**
    *   **Location:** A single flat directory (`/Notes/`) synced via Google Drive for Desktop.
    *   **Philosophy:** No subfolders. Organization relies on consistent naming conventions (lowercase, dashed slugs) and metadata rather than folder hierarchy.

*   **The Workflows:**
    *   **Keep to Markdown Migration:**
        *   A specific Gemini CLI command will extract data from Google Keep.
        *   It transforms these raw notes into structured Markdown files in the `/Notes/` directory.
        *   It archives the original Keep note to keep the inbox clean.

    *   **Daily Note Processing (TBD):**
        *   *Concept:* A workflow to process a single "Daily Note" file where random thoughts and meetings are logged throughout the day.
        *   *Refactoring:* AI will analyze the daily note and extract distinct topics into their own permanent files (e.g., moving a meeting summary to `project-alpha-meeting.md`).
        *   *Task Management:* AI will scan the daily note for action items (e.g., `- [ ]`) and aggregate them into a central `todo.md` or task tracking system.

### 3. Retrieving Notes

Retrieval leverages semantic understanding over simple text matching.

*   **Desktop:**
    *   **Tool:** Gemini CLI.
    *   **Workflow:** Run semantic searches directly in the terminal.
    *   **Example:** "What were the constraints we decided on for the 2026 API refactor?" (Gemini scans the `/Notes/` directory to synthesize an answer).

*   **Mobile:**
    *   **Tool:** Gemini App (iOS/Android).
    *   **Workflow:** Use the **@Google Drive** extension within the Gemini app to "talk" to your markdown files.
    *   **Example:** "Summarize my notes from yesterday's meeting."

---

## Architecture: The Notes CLI

To ensure robustness, portability, and "command-line first" capability, the core logic will be implemented as a standalone CLI tool written in **Go**.

*   **Project Name:** `notes-cli` (tentative)
*   **Language:** Go (Golang)
*   **Purpose:** Serve as the backend engine for the notes system. It handles file operations, parsing, and potential API interactions.
*   **Key Responsibilities:**
    *   Creating new notes with standardized frontmatter.
    *   Managing the Daily Note (creation, retrieval).
    *   Parsing and aggregating tasks.
    *   Interfacing with the Gemini CLI for AI tasks (or calling LLMs directly if needed later).
*   **Integration:** Neovim (and other tools) will interface with this CLI via shell commands.

---

## Implementation Roadmap

### Phase 1: Core Tooling (The Backend)
- [ ] **Notes CLI (Go Project):**
    -   Initialize a new Go project.
    -   Implement `notes new "<Title>"`:
        -   Generates slug.
        -   Creates file with YAML template.
        -   Outputs path for editor consumption.
- [ ] **Tagging Autocomplete:**
    -   Configure LazyVim and `blink.cmp` to enable autocomplete for tags within the YAML frontmatter of notes.

### Phase 2: Workflow Definitions
- [ ] **Define Daily Note Workflow:**
    -   Document the structure and usage of the daily note (e.g., capture format, timestamping).

### Phase 3: AI Agents & Automation
- [ ] **Daily Note Processor:**
    -   Develop an AI workflow to analyze the daily note, identify distinct topics, and refactor them into permanent, organized notes.
- [ ] **Task Manager:**
    -   Develop an AI workflow to extract action items from daily notes and organize them into a central task list.
- [ ] **Keep Importer:**
    -   Create the process/script to fetch notes from Google Keep, convert them to Markdown, and place them in the `/Notes/` directory.

### Phase 4: Neovim Integration
- [ ] **Workflow Plugin:**
    -   Develop a custom Neovim plugin to orchestrate the system.
    -   Features:
        -   Interface for rapid new note creation.
        -   Commands to trigger AI organization workflows (Keep sync, Daily refactor, Task extraction) directly from the editor.

