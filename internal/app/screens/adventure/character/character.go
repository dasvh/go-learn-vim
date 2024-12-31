package character

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/dasvh/go-learn-vim/internal/style"
)

// Character in the app
type Character struct {
	Rune   rune
	String string
}

// Characters represents the app characters
type Characters struct {
	Player struct {
		Cursor Character
		Trail  Character
	}
	Target struct {
		Active   Character
		Inactive Character
		Reached  Character
	}
}

// Style for the app characters
type Style struct {
	Player style.Player
	Target style.Target
}

// DefaultCharacters represents the default characters in the app
var DefaultCharacters = Characters{
	Player: struct {
		Cursor Character
		Trail  Character
	}{
		Cursor: Character{'$', "$"},
		Trail:  Character{'·', "·"},
	},
	Target: struct {
		Active   Character
		Inactive Character
		Reached  Character
	}{
		Active:   Character{'X', "X"},
		Inactive: Character{'x', "x"},
		Reached:  Character{'✓', "✓"},
	},
}

// ToDefaultCharacterStyle maps a rune to a default character style
func ToDefaultCharacterStyle(r rune) lipgloss.Style {
	switch r {
	case DefaultCharacters.Player.Cursor.Rune:
		return style.Styles.Adventure.Map.Player.Cursor
	case DefaultCharacters.Player.Trail.Rune:
		return style.Styles.Adventure.Map.Player.Trail
	case DefaultCharacters.Target.Active.Rune:
		return style.Styles.Adventure.Map.Target.Active
	case DefaultCharacters.Target.Inactive.Rune:
		return style.Styles.Adventure.Map.Target.Inactive
	case DefaultCharacters.Target.Reached.Rune:
		return style.Styles.Adventure.Map.Target.Reached
	default:
		return style.Styles.Adventure.Map.Background
	}
}
