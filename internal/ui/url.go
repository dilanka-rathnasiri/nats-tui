package ui

import (
	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
)

type urlModel struct {
	txtInput textinput.Model
	natsUrl  *string
}

func newUrlModel(natsUrl *string) urlModel {
	txt := textinput.New()
	txt.Focus()

	return urlModel{
		txtInput: txt,
		natsUrl:  natsUrl,
	}
}

func (m urlModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m urlModel) Update(msg tea.Msg) (urlModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "enter":
			*m.natsUrl = m.txtInput.Value()
			return m, nil
		}
	}

	var cmd tea.Cmd
	m.txtInput, cmd = m.txtInput.Update(msg)
	return m, cmd
}

func (m urlModel) View() tea.View {
	s := "Enter the nats server url?\n"
	s += m.txtInput.View()
	s += "\nPress enter to submit."
	s += "natsUrl: " + *m.natsUrl

	v := tea.NewView(s)
	v.AltScreen = true
	return v
}
