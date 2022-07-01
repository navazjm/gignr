package tui

import (
	"github.com/charmbracelet/bubbles/key"
)

type listKeyMap struct {
	generateGitignore    key.Binding
	deselectAllTemplates key.Binding
}

func newListKeyMap() *listKeyMap {
	return &listKeyMap{
		generateGitignore: key.NewBinding(
			key.WithKeys("ctrl+g"),
			key.WithHelp("ctrl+g", "generate"),
		),
		deselectAllTemplates: key.NewBinding(
			key.WithKeys("ctrl+d"),
			key.WithHelp("ctrl+d", "deselect all"),
		),
	}
}
