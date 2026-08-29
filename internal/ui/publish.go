package ui

import (
	tea "charm.land/bubbletea/v2"
)

type publishModel struct {
}

func newPublishModel() publishModel {
	return publishModel{}
}

func (m publishModel) Init() tea.Cmd {
	return nil
}

func (m publishModel) Update(msg tea.Msg) (publishModel, tea.Cmd) {
	return m, nil
}

func (m publishModel) View() tea.View {
	s := "You are in publish View"

	v := tea.NewView(s)
	v.AltScreen = true
	return v
}
