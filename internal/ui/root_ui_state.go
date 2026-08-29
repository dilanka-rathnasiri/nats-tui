package ui

type RootUiState int

const (
	initRootState RootUiState = iota
	urlRootState
	publishRootState
	subscribeRootState
	requestRootState
	replyRootState
)
