package style

import "github.com/charmbracelet/lipgloss"

// colours defines the color palette
var colours = struct {
	GoBlue     lipgloss.Color
	DarkBlue   lipgloss.Color
	Pink       lipgloss.Color
	LightPink  lipgloss.Color
	DarkPink   lipgloss.Color
	Green      lipgloss.Color
	LightGreen lipgloss.Color
	White      lipgloss.Color
	Grey       lipgloss.Color
	Black      lipgloss.Color
}{
	GoBlue:     lipgloss.Color("#00ADD8"),
	DarkBlue:   lipgloss.Color("#113344"),
	Pink:       lipgloss.Color("200"),
	LightPink:  lipgloss.Color("219"),
	DarkPink:   lipgloss.Color("125"),
	Green:      lipgloss.Color("120"),
	LightGreen: lipgloss.Color("157"),
	White:      lipgloss.Color("15"),
	Grey:       lipgloss.Color("236"),
	Black:      lipgloss.Color("16"),
}

// theme defines the color palette for the application
var theme = struct {
	Primary   lipgloss.Color
	Secondary lipgloss.Color
	Content   lipgloss.Color
	BgLight   lipgloss.Color
	BgDark    lipgloss.Color
}{
	Primary:   colours.GoBlue,
	Secondary: colours.DarkBlue,
	Content:   colours.Green,
	BgLight:   colours.White,
	BgDark:    colours.Pink,
}

// adventureTheme defines the color palette for the adventure game
var adventureTheme = struct {
	HeaderBg        lipgloss.Color
	FieldBg         lipgloss.Color
	HeaderBorderBg  lipgloss.Color
	HeaderBorderFg  lipgloss.Color
	FieldBorderBg   lipgloss.Color
	FieldBorderFg   lipgloss.Color
	PlayerFg        lipgloss.Color
	PlayerBg        lipgloss.Color
	TrailFg         lipgloss.Color
	TargetFg        lipgloss.Color
	TargetBg        lipgloss.Color
	TargetReachedFg lipgloss.Color
	TargetReachedBg lipgloss.Color
	WallFg          lipgloss.Color
	WallBg          lipgloss.Color
}{
	HeaderBg:        colours.DarkBlue,
	FieldBg:         colours.DarkBlue,
	HeaderBorderBg:  colours.DarkBlue,
	HeaderBorderFg:  colours.GoBlue,
	FieldBorderBg:   colours.GoBlue,
	FieldBorderFg:   colours.DarkBlue,
	PlayerFg:        colours.Pink,
	PlayerBg:        colours.Green,
	TrailFg:         colours.White,
	TargetFg:        colours.Pink,
	TargetBg:        colours.White,
	TargetReachedFg: colours.DarkBlue,
	TargetReachedBg: colours.Pink,
	WallFg:          colours.LightPink,
	WallBg:          colours.DarkPink,
}

// Target defines the styling for the target in the adventure game
type Target struct {
	Active   lipgloss.Style
	Inactive lipgloss.Style
	Reached  lipgloss.Style
}

// Player defines the styling for the player in the adventure game
type Player struct {
	Cursor lipgloss.Style
	Trail  lipgloss.Style
}

const itemIndicator = "•"
const selectedItemIndicator = "▸"

// Buttons defines the styling for buttons in the views components
var Buttons = struct {
	Layout   lipgloss.Position
	Selected lipgloss.Style
	Default  lipgloss.Style
}{
	Layout: lipgloss.Center,
	Selected: lipgloss.NewStyle().
		Width(30).
		Height(3).
		Border(lipgloss.NormalBorder()).
		Padding(1, 1).
		BorderForeground(colours.Pink).
		Background(colours.Pink).
		Foreground(colours.White),
	Default: lipgloss.NewStyle().
		Width(30).
		Height(3).
		Border(lipgloss.NormalBorder()).
		Padding(1, 1).
		BorderForeground(colours.Green).
		Background(colours.Green).
		Foreground(colours.Black),
}

// Table defines the styling for tables in the views components
var Table = struct {
	Header   lipgloss.Style
	Cell     lipgloss.Style
	Selected lipgloss.Style
}{
	Header: lipgloss.NewStyle().
		Foreground(theme.Primary).
		Background(colours.DarkBlue).
		Padding(0, 1).
		MarginBottom(1).
		Bold(true),
	Cell: lipgloss.NewStyle().
		Padding(0, 1),
	Selected: lipgloss.NewStyle().
		Foreground(colours.LightGreen).
		Background(colours.DarkPink).
		Bold(true),
}

// PlayerSelection defines the styling for player selection
var PlayerSelection = struct {
	Title        lipgloss.Style
	Item         lipgloss.Style
	SelectedItem lipgloss.Style
	FilterMatch  lipgloss.Style
	DimmedItem   lipgloss.Style
}{
	Title: lipgloss.NewStyle().
		Foreground(theme.Primary).
		Background(colours.DarkBlue).
		Width(40).
		Align(lipgloss.Center).
		Bold(true),
	Item: lipgloss.NewStyle().
		Foreground(colours.Pink).
		BorderLeft(true).
		BorderStyle(lipgloss.Border{Left: itemIndicator}).
		BorderForeground(colours.Pink).
		MarginLeft(2).
		PaddingLeft(1),
	SelectedItem: lipgloss.NewStyle().
		Foreground(colours.Green).
		BorderLeft(true).
		BorderStyle(lipgloss.Border{Left: selectedItemIndicator}).
		BorderForeground(colours.Green).
		MarginLeft(2).
		PaddingLeft(2).
		Bold(true),
	FilterMatch: lipgloss.NewStyle().
		Foreground(colours.Green),
	DimmedItem: lipgloss.NewStyle().
		Foreground(colours.Grey),
}

// LevelSelection defines the styling for level selection
var LevelSelection = struct {
	Title           lipgloss.Style
	Item            lipgloss.Style
	SelectedItem    lipgloss.Style
	Details         lipgloss.Style
	SelectedDetails lipgloss.Style
	FilterMatch     lipgloss.Style
	DimmedItem      lipgloss.Style
}{
	Title: lipgloss.NewStyle().
		Foreground(theme.Primary).
		Background(colours.DarkBlue).
		Width(40).
		Align(lipgloss.Center).
		Bold(true),
	Item: lipgloss.NewStyle().
		Foreground(colours.Pink).
		BorderLeft(true).
		BorderStyle(lipgloss.Border{Left: itemIndicator}).
		BorderForeground(colours.Pink).
		MarginLeft(2).
		PaddingLeft(1),
	SelectedItem: lipgloss.NewStyle().
		Foreground(colours.Green).
		BorderLeft(true).
		BorderStyle(lipgloss.Border{Left: selectedItemIndicator}).
		BorderForeground(colours.Green).
		MarginLeft(2).
		PaddingLeft(2).
		Bold(true),
	Details: lipgloss.NewStyle().
		Foreground(colours.LightPink).
		MarginLeft(3).
		PaddingLeft(1),
	SelectedDetails: lipgloss.NewStyle().
		Foreground(colours.LightGreen).
		MarginLeft(3).
		PaddingLeft(2).
		Bold(true),
	FilterMatch: lipgloss.NewStyle().
		Align(lipgloss.Center).
		Foreground(colours.Green),
	DimmedItem: lipgloss.NewStyle().
		Foreground(colours.Grey),
}

// Styles defines the styling configuration for views components
var Styles = struct {
	Title    lipgloss.Style
	Subtitle lipgloss.Style
	Display  struct {
		Section     lipgloss.Style
		Title       lipgloss.Style
		Text        lipgloss.Style
		SpecialText lipgloss.Style
	}
	Adventure struct {
		Header struct {
			Level  lipgloss.Style
			Mode   lipgloss.Style
			Stats  lipgloss.Style
			Border lipgloss.Style
		}
		Instructions struct {
			Style lipgloss.Style
		}
		Map struct {
			Border     lipgloss.Style
			Background lipgloss.Style
			Player     Player
			Target     Target
			Wall       lipgloss.Style
		}
	}
}{
	Title: lipgloss.NewStyle().
		Foreground(theme.Primary).
		Bold(true),
	Subtitle: lipgloss.NewStyle().
		Foreground(theme.Secondary).
		Background(theme.BgLight).
		Height(1).
		Width(45).
		MarginBottom(1).
		Bold(true).
		Blink(true).
		Align(lipgloss.Center),
	Display: struct {
		Section     lipgloss.Style
		Title       lipgloss.Style
		Text        lipgloss.Style
		SpecialText lipgloss.Style
	}{
		Section: lipgloss.NewStyle().
			BorderForeground(theme.Content).
			Border(lipgloss.NormalBorder()).
			Padding(1, 1),
		Title: lipgloss.NewStyle().
			Foreground(theme.Secondary).
			Background(theme.BgDark).
			MarginBottom(1).
			Bold(true).
			Underline(true),
		Text: lipgloss.NewStyle().
			Foreground(theme.Content).
			Padding(0, 1),
		SpecialText: lipgloss.NewStyle().
			Foreground(theme.Primary).
			Bold(true).
			MarginTop(1).
			Align(lipgloss.Center),
	},
	Adventure: struct {
		Header struct {
			Level  lipgloss.Style
			Mode   lipgloss.Style
			Stats  lipgloss.Style
			Border lipgloss.Style
		}
		Instructions struct {
			Style lipgloss.Style
		}
		Map struct {
			Border     lipgloss.Style
			Background lipgloss.Style
			Player     Player
			Target     Target
			Wall       lipgloss.Style
		}
	}{
		Header: struct {
			Level  lipgloss.Style
			Mode   lipgloss.Style
			Stats  lipgloss.Style
			Border lipgloss.Style
		}{
			Level: lipgloss.NewStyle().
				Foreground(colours.Green).
				Bold(true),
			Mode: lipgloss.NewStyle().
				Foreground(colours.Pink).
				Bold(true),
			Stats: lipgloss.NewStyle().
				Foreground(colours.Pink).
				Bold(true),
			Border: lipgloss.NewStyle().
				BorderStyle(lipgloss.ThickBorder()).
				PaddingLeft(1).
				PaddingRight(1).
				BorderForeground(adventureTheme.HeaderBorderFg).
				BorderBackground(adventureTheme.HeaderBorderBg),
		},
		Instructions: struct {
			Style lipgloss.Style
		}{
			Style: lipgloss.NewStyle().
				Foreground(colours.Green).
				Bold(true),
		},
		Map: struct {
			Border     lipgloss.Style
			Background lipgloss.Style
			Player     Player
			Target     Target
			Wall       lipgloss.Style
		}{
			Border: lipgloss.NewStyle().
				BorderStyle(lipgloss.RoundedBorder()).
				BorderForeground(adventureTheme.HeaderBorderFg).
				BorderBackground(adventureTheme.HeaderBorderBg),
			Background: lipgloss.NewStyle().
				Background(adventureTheme.FieldBg),
			Player: Player{
				Cursor: lipgloss.NewStyle().
					Foreground(adventureTheme.PlayerFg).
					Background(adventureTheme.PlayerBg).
					Bold(true).
					Blink(true),
				Trail: lipgloss.NewStyle().
					Foreground(adventureTheme.TrailFg).
					Background(colours.DarkBlue).
					Bold(true),
			},
			Target: Target{
				Active: lipgloss.NewStyle().
					Foreground(adventureTheme.TargetFg).
					Background(adventureTheme.TargetBg).
					Bold(true).
					Blink(true),
				Inactive: lipgloss.NewStyle().
					Foreground(adventureTheme.TargetFg),
				Reached: lipgloss.NewStyle().
					Foreground(adventureTheme.TargetReachedFg).
					Background(adventureTheme.TargetReachedBg).
					Bold(true).
					Blink(true),
			},
			Wall: lipgloss.NewStyle().
				Foreground(adventureTheme.WallFg).
				Background(adventureTheme.WallBg),
		},
	},
}

// GetComponentWidth calculates the total space taken by borders and padding in a lipgloss.Style
// by comparing the width of a rendered test string with its original width
func GetComponentWidth(style lipgloss.Style) int {
	testStr := "x"
	rendered := style.Render(testStr)
	return lipgloss.Width(rendered) - lipgloss.Width(testStr)
}

// GetComponentHeight calculates the total space taken by a component in a lipgloss.Style
func GetComponentHeight(style lipgloss.Style) int {
	testStr := "x"
	rendered := style.Render(testStr)
	return lipgloss.Height(rendered)
}
