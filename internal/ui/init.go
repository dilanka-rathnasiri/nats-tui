package ui

import (
	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
)

type initModel struct {
	txtInput textinput.Model
	natsUrl  string
}

func newInitModel() initModel {
	txt := textinput.New()
	txt.Focus()

	return initModel{
		txtInput: txt,
	}
}

func (m initModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m initModel) Update(msg tea.Msg) (initModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "enter":
			m.natsUrl = m.txtInput.Value()
			return m, nil
		}
	}

	var cmd tea.Cmd
	m.txtInput, cmd = m.txtInput.Update(msg)
	return m, cmd
}

func (m initModel) View() tea.View {
	s := "Enter the nats server url?\n"
	s += m.txtInput.View()
	s += "\nPress enter to submit."
	s += "\nPress ctrl+c to quit.\n"
	s += "natsUrl: " + m.natsUrl

	v := tea.NewView(s)
	v.AltScreen = true
	return v
}
