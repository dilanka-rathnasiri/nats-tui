package ui

import (
	tea "charm.land/bubbletea/v2"
)

type replyModel struct {
	natsUrl *string
}

func newReplyModel(natsUrl *string) replyModel {
	return replyModel{
		natsUrl: natsUrl,
	}
}

func (m replyModel) Init() tea.Cmd {
	return nil
}

func (m replyModel) Update(msg tea.Msg) (replyModel, tea.Cmd) {
	return m, nil
}

func (m replyModel) View() tea.View {
	s := "You are in reply View\n"
	s += "natsUrl: " + *m.natsUrl

	v := tea.NewView(s)
	v.AltScreen = true
	return v
}
