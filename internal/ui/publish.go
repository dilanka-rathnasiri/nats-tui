package ui

import (
	tea "charm.land/bubbletea/v2"
)

type publishModel struct {
	natsUrl *string
}

func newPublishModel(natsUrl *string) publishModel {
	return publishModel{
		natsUrl: natsUrl,
	}
}

func (m publishModel) Init() tea.Cmd {
	return nil
}

func (m publishModel) Update(msg tea.Msg) (publishModel, tea.Cmd) {
	return m, nil
}

func (m publishModel) View() tea.View {
	s := "You are in publish View\n"
	s += "natsUrl: " + *m.natsUrl

	v := tea.NewView(s)
	v.AltScreen = true
	return v
}
