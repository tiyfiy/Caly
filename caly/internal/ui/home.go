package ui

import (
	"github.com/charmbracelet/bubbles/list"
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
	//classes    []class
	cursor     int
	statusLine string
	width      int
	height     int
}

func initialModel() model {
	return model{
		cursor:     0,
		statusLine: "press q to quit",
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
	return m, nil
}

func (m model) View() string {
	header := headerStyle.Render("caly")

	content := contentStyle.Render("no events yet.")

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
