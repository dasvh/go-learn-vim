package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dasvh/go-learn-vim/internal/game"
	"os"
)

func main() {
	program := tea.NewProgram(game.NewGame())

	_, err := program.Run()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
