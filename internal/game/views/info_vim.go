package views

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/dasvh/go-learn-vim/internal/game/state"
	ui "github.com/dasvh/go-learn-vim/internal/ui/menu"
)

// Section represents a section in the Vim Information screen
type Section struct {
	Title   string
	Content string
}

// InfoVim represents the Vim Information screen
type InfoVim struct {
	*ui.BaseMenu
}

// NewInfoVimView returns a new InfoVim instance
func NewInfoVimView() ui.Menu {
	base := ui.NewBaseMenu("Vim Information", nil)
	content := renderSections(vimInfoSections)
	base.SetContent(content)
	return &InfoVim{BaseMenu: base}
}

// Update handles messages and transitions back to the Info menu
func (m *InfoVim) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	_, cmd := m.BaseMenu.Update(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if key.Matches(msg, m.Controls().Back) {
			return m, state.ChangeView(state.InfoView)
		}
	}

	return m, cmd
}

// HandleSelection implements ButtonHandler interface
func (m *InfoVim) HandleSelection() tea.Cmd { return nil }

func renderSections(sections []Section) string {
	var rendered string
	for _, section := range sections {
		rendered += lipgloss.JoinVertical(0,
			ui.ContentTitleStyle.Render(section.Title),
			ui.ContentBodyStyle.Render(section.Content),
		)
		rendered += "\n\n"
	}
	return ui.ContentSectionStyle.Render(rendered)
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
		Content: `- **Modal Editing**: Switch between modes for navigating, editing, and selecting text.
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
}
