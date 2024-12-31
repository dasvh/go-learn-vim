package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dasvh/go-learn-vim/internal/game"
	"github.com/dasvh/go-learn-vim/internal/storage"
	"os"
)

func main() {
	repo := storage.NewJSONRepository("adventure.json")
	program := tea.NewProgram(game.NewGame(repo))

	_, err := program.Run()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
