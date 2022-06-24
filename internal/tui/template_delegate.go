package tui

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type templateItemDelegate struct {
	keys *delegateKeyMap
}

func (d templateItemDelegate) Height() int  { return 1 }
func (d templateItemDelegate) Spacing() int { return 0 }

func (d templateItemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd {
	t, ok := m.SelectedItem().(templateItem)
	if !ok {
		return nil
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, d.keys.pick):
			t.isSelected = !t.isSelected
			return m.SetItem(m.Index(), t)
		}
	}

	return nil
}

func (d templateItemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	t, ok := listItem.(templateItem)
	if !ok {
		return
	}

	fn := lipgloss.NewStyle().PaddingLeft(4).Foreground(lipgloss.AdaptiveColor{Light: "#3D3D3D", Dark: "#FFFDF5"}).Render

	if t.IsSelected() {
		fn = func(s string) string {
			return lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.AdaptiveColor{Light: "#4078F2", Dark: "#57A5E5"}).Render("  " + s)
		}
	}

	if index == m.Index() {
		fn = func(s string) string {
			return lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.AdaptiveColor{Light: "#A626A4", Dark: "#BB70D2"}).Render("  " + s)
		}
	}

	fmt.Fprint(w, fn(t.title))
}

func (d templateItemDelegate) ShortHelp() []key.Binding {
	help := []key.Binding{d.keys.pick}
	return help
}

func (d templateItemDelegate) FullHelp() [][]key.Binding {
	help := []key.Binding{d.keys.pick}
	return [][]key.Binding{help}
}

func newTemplateItemDelegate(keys *delegateKeyMap) templateItemDelegate {
	return templateItemDelegate{
		keys: keys,
	}
}

type delegateKeyMap struct {
	pick key.Binding
}

// Additional short help entries. This satisfies the help.KeyMap interface and
// is entirely optional.
func (d delegateKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		d.pick,
	}
}

// Additional full help entries. This satisfies the help.KeyMap interface and
// is entirely optional.
func (d delegateKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{
			d.pick,
		},
	}
}

func newDelegateKeyMap() *delegateKeyMap {
	return &delegateKeyMap{
		pick: key.NewBinding(
			key.WithKeys(" "),
			key.WithHelp("space", "select/deselect"),
		),
	}
}
