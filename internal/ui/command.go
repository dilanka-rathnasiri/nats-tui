package ui

import (
	tea "charm.land/bubbletea/v2"
)

type commandModel struct {
}

func newCommandModel() commandModel {
	return commandModel{}
}

func (m commandModel) Init() tea.Cmd {
	return nil
}

func (m commandModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m commandModel) View() tea.View {
	s := "You are in command View"

	v := tea.NewView(s)
	v.AltScreen = true
	return v
}
