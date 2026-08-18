# TUI Template

This is a template for building complex Terminal User Interfaces (TUI) in Golang with a pre-configured architecture it includes a flexible design system that adapts to different screen sizes, a central place for managing reusable components and a system for handling both persistent and temporary modal panels.

## Highlights

*   Flexible grid system: Automatically switches between stacked layouts based on terminal width.
*   Panel Architecture: blocks (Lists, Tables, Trees, Inputs) that handle their own internal state and rendering.
*   Modal System: Support for centered overlays including alerts, confirmation dialogs, fuzzy selectors and structured forms.
*   Safe Rendering: Built-in height clamping and scrollbar logic to prevent content from breaking borders or overflowing the terminal.
*   Keyboard Management: keymaps and focus-tracking to prevent global shortcuts from interfering with text input.

## Overview

This template acts as a block showcase. The sidebar contains a catalog of UI components. Selecting a panel entry previews it in the workspace while selecting a modal entry triggers an interactive overlay.

The project is structured into two internal packages:
*   internal/ui/panels: Bordered components that occupy space in the layout grid.
*   internal/ui/modals: Focused overlays that exist outside the grid layout.

## Getting Started

### Prerequisites

*   Go 1.23.0 or higher

### Running the Showcase

To run the interactive component library and see the layout in action:

```bash
go run cmd/app/main.go
```

### Key Bindings

*   Tab / Shift+Tab: Cycle focus between the Sidebar and the Preview panel.
*   Enter: Activate the selected block or modal.
*   /: Filter the sidebar list.
*   q: Quit the application (suppressed when typing in text fields).
*   Ctrl+C: Force quit.

## Layout Logic

The template uses a custom layout engine defined in `internal/ui/layout.go`. This calculates the available screen space by accounting for the Titlebar, Statusbar, and Command Bar. 

When the terminal is less than 70 characters wide, the sidebar is placed above the main panel. When the terminal is wider than 70 characters, the sidebar and main panel are displayed side by side.

## Credits

This template is built upon the ecosystem provided by Charmbracelet. Credits to the following libraries:

*   [Bubble Tea: The functional TUI framework powering the application state.](https://github.com/charmbracelet/bubbletea)
*   [Lip Gloss: Used for all terminal styling, borders, and layout positioning.](https://github.com/charmbracelet/lipgloss)
*   [Bubbles: Provides the underlying logic for the lists, text inputs, viewports, and progress bars.](https://github.com/charmbracelet/bubbles)
*   [Huh: Powers the structured form modals.](https://github.com/charmbracelet/huh)
