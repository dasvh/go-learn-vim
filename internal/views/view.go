package views

import (
	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dasvh/go-learn-vim/internal/components"
	"github.com/dasvh/go-learn-vim/internal/style"
)

const mainTitle = `

 ██████╗  ██████╗     ██╗     ███████╗ █████╗ ██████╗ ███╗   ██╗    ██╗   ██╗██╗███╗   ███╗
██╔════╝ ██╔═══██╗    ██║     ██╔════╝██╔══██╗██╔══██╗████╗  ██║    ██║   ██║██║████╗ ████║
██║  ███╗██║   ██║    ██║     █████╗  ███████║██████╔╝██╔██╗ ██║    ██║   ██║██║██╔████╔██║
██║   ██║██║   ██║    ██║     ██╔══╝  ██╔══██║██╔══██╗██║╚██╗██║    ╚██╗ ██╔╝██║██║╚██╔╝██║
╚██████╔╝╚██████╔╝    ███████╗███████╗██║  ██║██║  ██║██║ ╚████║     ╚████╔╝ ██║██║ ╚═╝ ██║
 ╚═════╝  ╚═════╝     ╚══════╝╚══════╝╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═══╝      ╚═══╝  ╚═╝╚═╝     ╚═╝
`

// View represents a user interface component that adheres to the tea.Model interface
// and provides control mechanisms through the Controls method
type View interface {
	tea.Model
	Controls() components.Controls
}

// BaseView represents the base structure for a views component
// it contains the window size message, controls, help model, and header
type BaseView struct {
	size     tea.WindowSizeMsg
	controls components.Controls
	help     help.Model
	title    *components.Title
}

// NewBaseView creates a new instance of BaseView with the specified title
func NewBaseView(title string) *BaseView {
	return &BaseView{
		controls: components.NewControls(),
		help:     help.New(),
		title:    components.NewTitle(style.Styles.Title.Render(mainTitle), style.Styles.Subtitle.Render(title)),
	}
}

// Controls returns the Controls associated with the BaseView
func (bv *BaseView) Controls() components.Controls {
	return bv.controls
}
