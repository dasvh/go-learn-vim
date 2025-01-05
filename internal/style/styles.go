package style

import "github.com/charmbracelet/lipgloss"

var colours = struct {
	GoBlue     lipgloss.Color
	DarkBlue   lipgloss.Color
	Pink       lipgloss.Color
	Green      lipgloss.Color
	White      lipgloss.Color
	Grey       lipgloss.Color
	Red        lipgloss.Color
	Yellow     lipgloss.Color
	DarkGrey   lipgloss.Color
	LightPink  lipgloss.Color
	LightGreen lipgloss.Color
}{
	GoBlue:     lipgloss.Color("#00ADD8"),
	DarkBlue:   lipgloss.Color("#113344"),
	Pink:       lipgloss.Color("200"),
	Green:      lipgloss.Color("120"),
	White:      lipgloss.Color("15"),
	Grey:       lipgloss.Color("236"),
	Red:        lipgloss.Color("196"),
	Yellow:     lipgloss.Color("226"),
	DarkGrey:   lipgloss.Color("236"),
	LightPink:  lipgloss.Color("#FF1BA0"),
	LightGreen: lipgloss.Color("#A0FF1B"),
}

// theme defines the color palette for the views components
var theme = struct {
	Primary   lipgloss.Color
	Secondary lipgloss.Color
	Content   lipgloss.Color
	BgLight   lipgloss.Color
	BgDark    lipgloss.Color
}{
	Primary:   lipgloss.Color("#00ADD8"),
	Secondary: lipgloss.Color("#113344"),
	Content:   lipgloss.Color("120"),
	BgLight:   lipgloss.Color("#f0f0f0"),
	BgDark:    lipgloss.Color("200"),
}

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
	WallFg:          colours.LightGreen,
	WallBg:          colours.LightPink,
}

type Target struct {
	Active   lipgloss.Style
	Inactive lipgloss.Style
	Reached  lipgloss.Style
}

type Player struct {
	Cursor lipgloss.Style
	Trail  lipgloss.Style
}

const itemIndicator = "•"
const selectedItemIndicator = "▸"

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
		Section lipgloss.Style
		Title   lipgloss.Style
		Text    lipgloss.Style
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
		Section lipgloss.Style
		Title   lipgloss.Style
		Text    lipgloss.Style
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
