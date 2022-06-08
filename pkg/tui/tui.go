package tui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/tehmj/gignr/pkg/utils"
)

type model struct {
	list      list.Model
	selection []string
}

func NewModel() model {

	gitignoreTemplates := utils.ConvertPathsToFilenames(utils.GetTemplates())

	var templates []list.Item

	templatesCount := len(gitignoreTemplates)
	for i := 0; i < templatesCount; i++ {
		newTemplateItem := templateItem{title: gitignoreTemplates[i]}
		templates = append(templates, newTemplateItem)
	}

	templateList := list.New(templates, templateItemDelegate{}, 0, 0)
	templateList.Title = "Welcome to Gignr - Generate .gitignore templates at ease"
	templateList.Styles.Title = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFDF5"))
	templateList.Paginator.PerPage = 25
	templateList.Paginator.SetTotalPages(len(gitignoreTemplates))

	return model{
		list: templateList,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
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
