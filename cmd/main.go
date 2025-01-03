package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dasvh/go-learn-vim/internal/app"
	"github.com/dasvh/go-learn-vim/internal/storage"
	"os"
)

func main() {
	repo, err := storage.NewJSONRepository("adventure.json")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	program := tea.NewProgram(app.NewApp(repo))

	_, err = program.Run()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
