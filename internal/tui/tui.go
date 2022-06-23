package tui

import (
	"bufio"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/paginator"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/navazjm/gignr/internal/utils"
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
		newTemplateItem := templateItem{title: gitignoreTemplates[i], isSelected: false}
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
		case m.keys.generateTemplate.Help().Key:
			// create list of selected templates
			var selectedTemplates []string
			for _, item := range m.list.Items() {
				templateItem := item.(templateItem)
				if templateItem.IsSelected() {
					selectedTemplates = append(selectedTemplates, templateItem.Title())
				}
			}

			if len(selectedTemplates) < 1 {
				statusCmd := m.list.NewStatusMessage(statusMessageStyle("Error! No templates selected"))
				return m, statusCmd
			}

			templatePaths := utils.GetTemplates()                             // used to read content of gitignore templates
			templateFilenames := utils.ConvertPathsToFilenames(templatePaths) // used to match specified templates by user

			var gitignoreContents []string // contents of the new gitignore file

			for _, template := range selectedTemplates {
				template = strings.ToLower(template)
				for i, file := range templateFilenames {
					if template != file {
						continue
					}

					templateFile, err := os.Open(templatePaths[i])

					if err != nil {
						statusCmd := m.list.NewStatusMessage(statusMessageStyle(err.Error()))
						return m, statusCmd
					}

					defer templateFile.Close()

					scanner := bufio.NewScanner(templateFile)
					for scanner.Scan() {
						line := scanner.Text()
						line += "\n"
						gitignoreContents = append(gitignoreContents, line)
					}

					if err := scanner.Err(); err != nil {
						statusCmd := m.list.NewStatusMessage(statusMessageStyle(err.Error()))
						return m, statusCmd
					}
				}
			}

			// get gitignore path in users current working dir
			gitignorePath, err := os.Getwd()
			if err != nil {
				statusCmd := m.list.NewStatusMessage(statusMessageStyle(err.Error()))
				return m, statusCmd
			}
			gitignorePath += "/.gitignore"

			var gitignoreFile *os.File

			gitignoreFile, err = os.Create(gitignorePath)

			if err != nil {
				statusCmd := m.list.NewStatusMessage(statusMessageStyle(err.Error()))
				return m, statusCmd
			}

			defer gitignoreFile.Close()

			for _, line := range gitignoreContents {
				if _, err := gitignoreFile.WriteString(line); err != nil {
					statusCmd := m.list.NewStatusMessage(statusMessageStyle(err.Error()))
					return m, statusCmd
				}
			}

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
