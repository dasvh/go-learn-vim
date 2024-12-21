package info

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/dasvh/go-learn-vim/internal/components/view"
	"github.com/dasvh/go-learn-vim/internal/game/state"
)

// VimInfo represents the Vim Information view
type VimInfo struct {
	*view.ContentView
}

// Section represents a section in the Vim Information view
type Section struct {
	Title   string
	Content string
}

// NewVimInfo creates and returns a new VimInfo instance with initialized display view
func NewVimInfo() *VimInfo {
	display := view.NewDisplayView("Vim Information")
	content := renderSections(vimInfoSections) + Bram
	display.SetContent(content)
	return &VimInfo{ContentView: display}
}

// Update handles updates to the VimInfo screen model based on incoming messages
func (vi *VimInfo) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	_, cmd := vi.ContentView.Update(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if key.Matches(msg, vi.Controls().Back) {
			return vi, state.ChangeScreen(state.InfoMenuScreen)
		}
	}

	return vi, cmd
}

// renderSections takes a slice of Section structs and returns a formatted string representation
func renderSections(sections []Section) string {
	var rendered string
	for _, section := range sections {
		rendered += lipgloss.JoinVertical(0,
			view.Styles.Display.Title.Render(section.Title),
			view.Styles.Display.Text.Render(section.Content),
		) + "\n\n"
	}
	return view.Styles.Display.Section.Render(rendered)
}

var vimInfoSections = []Section{
	{
		Title: "What is Vim?",
		Content: `Vim (Vi IMproved) is a highly efficient and feature-rich text editor designed for speed and productivity. 
Originally derived from the Vi editor in 1991, Vim builds on Vi’s simplicity while adding powerful features 
that make it a favorite tool for developers, system administrators, and text editing enthusiasts.`,
	},
	{
		Title: "Why Use Vim?",
		Content: `Vim is the default fallback editor on all POSIX systems. Whether you've just installed the operating system,
or you've booted into a minimal environment to repair a system, or you're unable to access any other editor,
Vim is sure to be available. While you can swap out other small editors, such as GNU Nano or Jove, on your system,
it's Vim that's all but guaranteed to be on every other system in the world.`,
	},
	{
		Title: "Key Features:",
		Content: `- **Modal Editing**: SwitchTo between modes for navigating, editing, and selecting text.
- **Customizability**: Personalize Vim with plugins, keybindings, and themes to suit your workflow.
- **Performance**: Designed to work in any environment, from a lightweight terminal to powerful plugins.
- **Availability**: Pre-installed on most Unix-based systems, making it accessible almost anywhere.`,
	},
	{
		Title: "Why Use Vim?",
		Content: `1. **Speed**: Optimized for keyboard-driven workflows, enabling faster text editing compared to GUI editors.
2. **Power**: Perform complex edits, automate repetitive tasks, and handle large files effortlessly.
3. **Portability**: Vim runs on virtually every platform, from servers to local machines, and is often 
   the go-to editor for remote editing over SSH.
4. **Extensibility**: Enhance functionality with plugins for linting, autocomplete, file navigation, and more.`,
	},
	{
		Title: "The Philosophy of Vim:",
		Content: `Vim embraces the philosophy of efficiency through practice. It has a learning curve, but as you grow familiar 
with its commands and capabilities, it becomes a tool that adapts to your needs, making your workflow faster 
and more seamless.`,
	},
	{
		Title: "Additional information:",
		Content: `For more information on Vim, visit the official website at https://www.vim.org/ 
For a comprehensive guide on Vim, check out the Vim documentation at https://vimhelp.org/`,
	},
}

const (
	Bram = `
Rest in peace, Bram Moolenaar, the creator of Vim, who passed away on August 3rd, 2023.
`
)
