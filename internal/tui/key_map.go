package tui

import (
	"github.com/charmbracelet/bubbles/key"
)

type listKeyMap struct {
	generateGitignore key.Binding
  deselectAllTemplates key.Binding
}

func newListKeyMap() *listKeyMap {
	return &listKeyMap{
		generateGitignore: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "generate"),
		),
		deselectAllTemplates: key.NewBinding(
			key.WithKeys("d"),
			key.WithHelp("d", "deselect all"),
		),
	}
}
