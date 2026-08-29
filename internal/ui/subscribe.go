package ui

import (
	tea "charm.land/bubbletea/v2"
)

type subscribeModel struct {
}

func newSubscribeModel() subscribeModel {
	return subscribeModel{}
}

func (m subscribeModel) Init() tea.Cmd {
	return nil
}

func (m subscribeModel) Update(msg tea.Msg) (subscribeModel, tea.Cmd) {
	return m, nil
}

func (m subscribeModel) View() tea.View {
	s := "You are in subscribe View"

	v := tea.NewView(s)
	v.AltScreen = true
	return v
}
