package ui

import (
	tea "charm.land/bubbletea/v2"
)

type requestModel struct {
	natsUrl *string
}

func newRequestModel(natsUrl *string) requestModel {
	return requestModel{
		natsUrl: natsUrl,
	}
}

func (m requestModel) Init() tea.Cmd {
	return nil
}

func (m requestModel) Update(msg tea.Msg) (requestModel, tea.Cmd) {
	return m, nil
}

func (m requestModel) View() tea.View {
	s := "You are in request View\n"
	s += "natsUrl: " + *m.natsUrl

	v := tea.NewView(s)
	v.AltScreen = true
	return v
}
