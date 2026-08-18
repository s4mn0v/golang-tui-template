<p align="center">
  <img width="400" alt="image" src="https://github.com/user-attachments/assets/2df5ce4c-9ebd-4c92-a21c-f8334035c848" />
  <img width="400" alt="image" src="https://github.com/user-attachments/assets/9a080c12-eb7e-47c5-b0c3-2da047be0a8c" />
</p>


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

## Preview
> [!NOTE]
> This is just screenshots of some of the blocks/highlights of the template to see the full list of the block check the drop down bellow
<details>
<summary>FULL LIST BLOCKS</summary>

## Data / Lists

- `list` — Navigable list with fuzzy filtering
- `table` — Table with rows and columns
- `viewport` — Scrollable long-form text (logs, results, docs)
- `paginator` — Pagination indicator/control

## Input

- `textinput` — Single-line input field
- `textarea` — Multiline input field
- `filepicker` — File/directory picker

## State / Feedback

- `spinner` — Loading indicator
- `progress` — Progress bar
- `help` — Help panel with keyboard shortcuts
- `key` — Key/shortcut indicator

## Structure

- `statusbar` — Status bar (footer)
- `titlebar` — Title bar (header)
- `tabs` — Tab navigation
- `tree` — Hierarchical tree (schemas, filesystem, git)

## Modals (overlay, not part of the grid)

- `modal:alert` — Info / error / success message with one button
- `modal:confirm` — Yes/no confirmation
- `modal:form` — Form
- `modal:selector` — Picker with fuzzy search

## Generic

- `custom` — Empty block with no assigned component (placeholder for custom logic)

</details>
<img width="1919" height="1005" alt="image" src="https://github.com/user-attachments/assets/58353d9b-0be6-421c-bd24-5c48689ce319" />
<img width="1919" height="1005" alt="image" src="https://github.com/user-attachments/assets/cd33e37a-9b0a-48a0-af7f-4d3be8044fcb" />
<img width="1919" height="1005" alt="image" src="https://github.com/user-attachments/assets/e59e3cb7-a0eb-4a21-8acc-e02215bd28c5" />
<img width="1919" height="1005" alt="image" src="https://github.com/user-attachments/assets/026c0210-dca5-4b06-81f1-bc02f1ce6208" />
<img width="1919" height="1005" alt="image" src="https://github.com/user-attachments/assets/46b6b25a-2335-4039-baeb-4d62a80cc7df" />


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
