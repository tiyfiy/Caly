package ui

import (
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/tiyfiy/caly/internal/data"
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

type model struct {
	table      table.Model
	hours      []data.Hour
	cursor     int
	statusLine string
	width      int
	height     int
}

func initialModel() model {
	return model{
		table:      newTable(nil, 0),
		cursor:     0,
		statusLine: "loading...",
	}
}

func (m model) Init() tea.Cmd {
	return data.FetchHours()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case data.HoursFetchedMsg:
		m.hours = msg.Hours
		m.table = newTable(m.hours, m.height)
		m.statusLine = "↑/↓ navigate • enter select • q quit"

	case data.HoursErrMsg:
		m.statusLine = "error: " + msg.Err.Error()

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.table = newTable(m.hours, m.height)

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
