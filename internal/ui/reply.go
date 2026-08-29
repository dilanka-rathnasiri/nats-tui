package ui

import (
	tea "charm.land/bubbletea/v2"
)

type replyModel struct {
}

func newReplyModel() replyModel {
	return replyModel{}
}

func (m replyModel) Init() tea.Cmd {
	return nil
}

func (m replyModel) Update(msg tea.Msg) (replyModel, tea.Cmd) {
	return m, nil
}

func (m replyModel) View() tea.View {
	s := "You are in reply View"

	v := tea.NewView(s)
	v.AltScreen = true
	return v
}
