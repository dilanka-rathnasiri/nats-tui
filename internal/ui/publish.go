package ui

import (
	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
)

type publishModel struct {
	natsUrl   *string
	txtInputs []textinput.Model
	focusIdx  int
}

func newPublishModel(natsUrl *string) publishModel {
	txtInputs := make([]textinput.Model, 2)

	for i := range 2 {
		txtInputs[i] = textinput.New()
	}
	txtInputs[0].Focus()

	return publishModel{
		natsUrl:   natsUrl,
		txtInputs: txtInputs,
		focusIdx:  0,
	}
}

func (m publishModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m publishModel) Update(msg tea.Msg) (publishModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "tab":
			if m.focusIdx == 0 {
				m.txtInputs[1].Focus()
				m.txtInputs[0].Blur()
				m.focusIdx = 1
			} else {
				m.txtInputs[0].Focus()
				m.txtInputs[1].Blur()
				m.focusIdx = 0
			}
		case "enter":
			return m, tea.Quit
		}
	}

	if m.focusIdx == 0 {
		var cmd tea.Cmd
		m.txtInputs[0], cmd = m.txtInputs[0].Update(msg)
		return m, cmd
	}
	var cmd tea.Cmd
	m.txtInputs[1], cmd = m.txtInputs[1].Update(msg)
	return m, cmd
}

func (m publishModel) View() tea.View {
	s := "Publish a message\n"
	s += "nats url: " + *m.natsUrl
	s += "\nnats subject: " + m.txtInputs[0].View()
	s += "\nnats message: " + m.txtInputs[1].View()
	s += "\nPress enter to submit."
	s += "\nPress tab to switch between text inputs."

	v := tea.NewView(s)
	v.AltScreen = true
	return v
}
