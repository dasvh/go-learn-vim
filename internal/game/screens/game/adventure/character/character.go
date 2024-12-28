package character

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/dasvh/go-learn-vim/internal/components/view"
)

// Character in the game
type Character struct {
	Rune   rune
	String string
}

// Characters represents the game characters
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

// Style for the game characters
type Style struct {
	Player view.Player
	Target view.Target
}

// DefaultCharacters represents the default characters in the game
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
		return view.Styles.Adventure.Map.Player.Cursor
	case DefaultCharacters.Player.Trail.Rune:
		return view.Styles.Adventure.Map.Player.Trail
	case DefaultCharacters.Target.Active.Rune:
		return view.Styles.Adventure.Map.Target.Active
	case DefaultCharacters.Target.Inactive.Rune:
		return view.Styles.Adventure.Map.Target.Inactive
	case DefaultCharacters.Target.Reached.Rune:
		return view.Styles.Adventure.Map.Target.Reached
	default:
		return view.Styles.Adventure.Map.Background
	}
}
