package ui

import (
	tea "charm.land/bubbletea/v2"
)

type rootModel struct {
	currntState  UiState
	initModel    initModel
	commandModel commandModel
}

func NewRootModel() rootModel {
	return rootModel{
		currntState:  initState,
		initModel:    newInitModel(),
		commandModel: newCommandModel(),
	}
}

func (m rootModel) Init() tea.Cmd {
	return nil
}

func (m rootModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		// global keybinding
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}

		// context aware key bindings
		switch m.currntState {
		case initState:
			var cmd tea.Cmd
			m.initModel, cmd = m.initModel.Update(msg) // delegate update to initModel

			switch msg.String() { // switch the state
			case "enter":
				m.currntState = commandState
			}

			return m, cmd
		case commandState:
			switch msg.String() {
			case "esc":
				m.currntState = initState
			}
		}
	}
	return m, nil
}

func (m rootModel) View() tea.View {
	switch m.currntState {
	case initState:
		return m.initModel.View()
	case commandState:
		return m.commandModel.View()
	}

	s := "You are in root View"
	v := tea.NewView(s)
	v.AltScreen = true
	return v
}
