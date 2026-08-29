package ui

import (
	tea "charm.land/bubbletea/v2"
)

type rootModel struct {
	currntState    RootUiState
	initModel      initModel
	urlModel       urlModel
	publishModel   publishModel
	subscribeModel subscribeModel
	requestModel   requestModel
	replyModel     replyModel
}

func NewRootModel() rootModel {
	return rootModel{
		currntState:    initRootState,
		initModel:      newInitModel(),
		urlModel:       newUrlModel(),
		publishModel:   newPublishModel(),
		subscribeModel: newSubscribeModel(),
		requestModel:   newRequestModel(),
		replyModel:     newReplyModel(),
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
		case "esc":
			m.currntState = initRootState
		case "ctrl+c":
			return m, tea.Quit
		}

		// context aware key bindings
		switch m.currntState {
		case initRootState:
			switch msg.String() {
			case "0":
				m.currntState = urlRootState
			case "1":
				m.currntState = publishRootState
			case "2":
				m.currntState = subscribeRootState
			case "3":
				m.currntState = requestRootState
			case "4":
				m.currntState = replyRootState
			}
		case urlRootState:
			var cmd tea.Cmd
			m.urlModel, cmd = m.urlModel.Update(msg) // delegate update to urlModel
			switch msg.String() {
			case "enter":
				m.currntState = initRootState
			}
			return m, cmd
		case publishRootState:
			var cmd tea.Cmd
			m.publishModel, cmd = m.publishModel.Update(msg) // delegate update to publishModel
			return m, cmd
		case subscribeRootState:
			var cmd tea.Cmd
			m.subscribeModel, cmd = m.subscribeModel.Update(msg) // delegate update to subscribeModel
			return m, cmd
		case requestRootState:
			var cmd tea.Cmd
			m.requestModel, cmd = m.requestModel.Update(msg) // delegate update to requestModel
			return m, cmd
		case replyRootState:
			var cmd tea.Cmd
			m.replyModel, cmd = m.replyModel.Update(msg) // delegate update to replyModel
			return m, cmd
		}
	}
	return m, nil
}

func (m rootModel) View() tea.View {
	switch m.currntState {
	case initRootState:
		return m.initModel.View() // delegate view to initModel
	case urlRootState:
		return m.urlModel.View() // delegate view to urlModel
	case publishRootState:
		return m.publishModel.View() // delegate view to publishModel
	case subscribeRootState:
		return m.subscribeModel.View() // delegate view to subscribeModel
	case requestRootState:
		return m.requestModel.View() // delegate view to requestModel
	case replyRootState:
		return m.replyModel.View() // delegate view to replyModel
	}

	s := "You are in root View"
	v := tea.NewView(s)
	v.AltScreen = true
	return v
}
