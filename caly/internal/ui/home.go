package ui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	headerStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("62")).
			Padding(0, 1)

	statusBarStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("241")).
			Padding(0, 1)

	contentStyle = lipgloss.NewStyle().
			Padding(1, 2)
)

type class struct {
	ID   string
	name string
}

type model struct {
	table      table.Model
	cursor     int
	statusLine string
	width      int
	height     int
}

func initialModel() model {
	return model{
		table:      newTable(),
		cursor:     0,
		statusLine: "↑/↓ navigate • enter select • q quit",
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m model) View() string {
	header := headerStyle.Render("caly")

	content := contentStyle.Render(m.table.View())

	statusBar := statusBarStyle.Render(m.statusLine)

	return lipgloss.JoinVertical(lipgloss.Left,
		header,
		content,
		statusBar,
	)
}

func Run() error {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	_, err := p.Run()
	return err
}

var _ list.Item = class{}

func (c class) Title() string       { return c.name }
func (c class) Description() string { return c.ID }
func (c class) FilterValue() string { return c.name }
