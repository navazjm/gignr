package tui

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/paginator"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/tehmj/gignr/pkg/utils"
)

var statusMessageStyle = lipgloss.NewStyle().
	Foreground(lipgloss.AdaptiveColor{Light: "#04B575", Dark: "#04B575"}).
	Render

type model struct {
	list         list.Model
	keys         *listKeyMap
	delegateKeys *delegateKeyMap
}

func NewModel() model {

	delegateKeys := newDelegateKeyMap()
	listKeys := newListKeyMap()

	gitignoreTemplates := utils.ConvertPathsToFilenames(utils.GetTemplates())

	var templates []list.Item

	templatesCount := len(gitignoreTemplates)
	for i := 0; i < templatesCount; i++ {
		newTemplateItem := templateItem{title: gitignoreTemplates[i]}
		templates = append(templates, newTemplateItem)
	}

	templateItemDelegate := newTemplateItemDelegate(delegateKeys)
	templateList := list.New(templates, templateItemDelegate, 0, 0)
	templateList.Title = "Gignr"
	templateList.Styles.Title = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFDF5"))
	templateList.Help.ShortSeparator = "  "
  templateList.Paginator.Type = paginator.Arabic
	templateList.AdditionalFullHelpKeys = func() []key.Binding {
		return []key.Binding{
			listKeys.generateTemplate,
		}
	}

	return model{
		list:         templateList,
		keys:         listKeys,
		delegateKeys: delegateKeys,
	}
}

func (m model) Init() tea.Cmd {
	return tea.EnterAltScreen
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := lipgloss.NewStyle().Padding(1, 2).GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd

}

func (m model) View() string {
	return lipgloss.NewStyle().Padding(1, 2).Render(m.list.View())
}
