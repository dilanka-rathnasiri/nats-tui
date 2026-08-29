package ui

import (
	tea "charm.land/bubbletea/v2"
)

type initModel struct {
}

func newInitModel() initModel {
	return initModel{}
}

func (m initModel) Init() tea.Cmd {
	return nil
}

func (m initModel) Update(msg tea.Msg) (initModel, tea.Cmd) {
	return m, nil
}

func (m initModel) View() tea.View {
	s := "0 => Url\n"
	s += "1 => Publish\n"
	s += "2 => Subscribe\n"
	s += "3 => Request\n"
	s += "4 => Reply\n"

	v := tea.NewView(s)
	v.AltScreen = true
	return v
}
