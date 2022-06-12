package tui

import (
	"github.com/charmbracelet/bubbles/key"
)

type listKeyMap struct {
	generateTemplate key.Binding
}

func newListKeyMap() *listKeyMap {
	return &listKeyMap{
		generateTemplate: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "generate"),
		),
	}
}
