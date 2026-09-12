package app

import (
	"log/slog"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/dilanka-rathnasiri/nats-tui/internal/ui"
)

func Run() {
	p := tea.NewProgram(ui.NewRootModel())
	if _, err := p.Run(); err != nil {
		slog.Error("error running program", "error", err)
		os.Exit(1)
	}
}
