package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/ssebs/tui-dash/tui"
)

func main() {
	fmt.Println("TUI Dash")
	// p := tea.NewProgram(tui.CreateExampleModel())
	p := tea.NewProgram(tui.CreateMainWindow())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
