package tui

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type templateItemDelegate struct{}

func (d templateItemDelegate) Height() int                               { return 1 }
func (d templateItemDelegate) Spacing() int                              { return 0 }
func (d templateItemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }
func (d templateItemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(templateItem)
	if !ok {
		return
	}

	str := fmt.Sprintf("%s", i)

	fn := lipgloss.NewStyle().PaddingLeft(4).Render

	if index == m.Index() {
		fn = func(s string) string {
			return lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170")).Render("> " + s)
		}
	}

	fmt.Fprintf(w, fn(str))
}
