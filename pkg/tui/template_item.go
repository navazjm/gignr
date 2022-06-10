package tui

type templateItem struct {
	title      string
	isSelected bool
}

func (t templateItem) Title() string       { return t.title }
func (t templateItem) FilterValue() string { return t.title }
func (t templateItem) IsSelected() bool    { return t.isSelected }
