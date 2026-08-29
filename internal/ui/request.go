package ui

import (
	tea "charm.land/bubbletea/v2"
)

type requestModel struct {
}

func newRequestModel() requestModel {
	return requestModel{}
}

func (m requestModel) Init() tea.Cmd {
	return nil
}

func (m requestModel) Update(msg tea.Msg) (requestModel, tea.Cmd) {
	return m, nil
}

func (m requestModel) View() tea.View {
	s := "You are in request View"

	v := tea.NewView(s)
	v.AltScreen = true
	return v
}
