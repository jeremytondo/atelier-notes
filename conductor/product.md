# Product Guide - Atelier Notes

## Initial Concept
Atelier Notes is a CLI tool designed to manage an AI-enhanced knowledge base. It facilitates note creation, organization, and retrieval using a flat file structure and semantic search capabilities.

## Target Audience
The primary audience for Atelier Notes consists of individual developers and power users. These users typically manage their notes in Markdown and prefer command-line interfaces for their speed and integration with other developer tools.

## Core Goals
*   **AI-Enhanced Intelligence:** Atelier Notes aims to transform a passive collection of text files into an active knowledge base by providing semantic search and AI-assisted organization, making information retrieval effortless.
*   **Workflow Friction Reduction:** The tool is designed to provide a fast CLI interface that bridges the gap between various capture methods, such as Neovim on desktop and Google Keep on mobile, and the local flat-file storage.

## Critical Features
*   **Advanced Retrieval:** Robust semantic search capabilities powered by the Gemini CLI to find notes based on meaning rather than just keywords.
*   **Seamless Capture:** High-fidelity integration between mobile capture (Google Keep), desktop editing (Neovim), and the local filesystem.
*   **Automated Organization:** AI-driven agents that automatically manage metadata and file structure, ensuring the knowledge base remains organized without manual effort.
*   **Natural Language Querying:** An `ask` command that allows users to query their notes using natural language, powered by the Gemini CLI.

## User Experience (UX) Philosophy
*   **Speed and Efficiency:** Atelier Notes prioritizes near-instantaneous, keyboard-centric operations to maintain user flow.
*   **Transparency:** All AI-driven modifications to notes, including metadata updates and file reorganization, must be visible to the user and easily reversible to ensure trust and control.

---

# Product Guide - Atelier-Notes.nvim

## Goal
The primary goal of this project is to provide a seamless, highly opinionated Neovim interface for interacting with the `atelier-notes` CLI. This plugin is specifically tailored to support a personal note-taking workflow and assumes a modern Neovim environment, particularly targeting the **LazyVim** distribution.

## Target Audience
- **Primary:** The author (personal use) and users with a similar, opinionated workflow.
- **Secondary:** Neovim users employing the LazyVim distribution who want a deeply integrated note-taking solution.

## Core Features
- **Note Creation:** Quickly create new notes from within Neovim using integrated prompts.
- **Search & Pick:** Seamless integration with standard LazyVim plugins (like `telescope.nvim` or `fzf-lua`) to find and open notes.
- **LazyVim Native:** Prioritizes leveraging features, utilities, and plugins already included in the LazyVim distribution to minimize external dependencies and ensure consistency.
- **CLI Wrapper:** A robust Lua interface that wraps `atelier-notes` commands, handling execution and result parsing.

## User Experience
The plugin is designed to fit perfectly into a LazyVim setup. It leverages the existing ecosystem's UI components and keybinding conventions to provide a fluid, "native" feel for managing notes without context switching.