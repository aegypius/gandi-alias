package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
)

func main() {
	log.Info("Starting application")
	p := tea.NewProgram(
		InitModel(),
	)
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}
