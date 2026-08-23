package ui

import (
	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
)

type initModel struct {
	txtInput textinput.Model
}

func InitModel() initModel {
	txt := textinput.New()
	txt.Focus()

	return initModel{
		txtInput: txt,
	}
}

func (m initModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m initModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.txtInput, cmd = m.txtInput.Update(msg)
	return m, cmd
}

func (m initModel) View() tea.View {
	s := "Enter the nats server url?\n"
	s += m.txtInput.View()
	s += "\nPress ctrl+c to quit.\n"

	v := tea.NewView(s)
	v.AltScreen = true
	return v
}
