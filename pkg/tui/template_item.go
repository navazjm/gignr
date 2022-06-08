package tui

type templateItem struct {
  title string
}

func (t templateItem) Title() string { return t.title }
func (t templateItem) FilterValue() string { return t.title }

