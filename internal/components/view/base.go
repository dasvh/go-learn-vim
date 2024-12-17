package view

import (
	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
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
	Controls() Controls
}

// header represents the structure for a header component with a main title and a subtitle.
type header struct {
	mainTitle string
	subtitle  string
}

// BaseView represents the base structure for a view component
// it contains the window size message, controls, help model, and header
type BaseView struct {
	size     tea.WindowSizeMsg
	controls Controls
	help     help.Model
	header
}

// NewBaseView creates a new instance of BaseView with the specified title
func NewBaseView(title string) *BaseView {
	return &BaseView{
		controls: NewControls(),
		help:     help.New(),
		header: header{
			mainTitle: Styles.Title.Render(mainTitle),
			subtitle:  Styles.Subtitle.Render(title),
		},
	}
}

// Controls returns the Controls associated with the BaseView
func (b *BaseView) Controls() Controls {
	return b.controls
}
