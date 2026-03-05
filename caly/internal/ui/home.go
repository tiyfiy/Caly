package ui

import (
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/tiyfiy/caly/internal/data"
)

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("62")).
			Padding(0, 1)

	weekStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("245")).
			Padding(0, 1)

	detailStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("252")).
			Padding(0, 2)

	detailLabelStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("255"))

	statusBarStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("241")).
			Padding(0, 1)
)

type model struct {
	hours      []data.Hour
	lectures   []data.Lecture
	grid       gridData
	cursorRow  int
	cursorCol  int
	weekOffset int
	weekStart  time.Time
	statusLine string
	width      int
	height     int
	loading    bool
}

func currentWeekStart() time.Time {
	now := time.Now()
	wd := int(now.Weekday())
	if wd == 0 {
		wd = 7
	}
	monday := now.AddDate(0, 0, -(wd - 1))
	return time.Date(monday.Year(), monday.Month(), monday.Day(), 0, 0, 0, 0, monday.Location())
}

func initialModel() model {
	ws := currentWeekStart()
	return model{
		weekStart:  ws,
		statusLine: "loading...",
		loading:    true,
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(data.FetchHours(), data.FetchLectures())
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case data.HoursFetchedMsg:
		m.hours = msg.Hours
		m.rebuildGrid()
		m.loading = false
		m.statusLine = "↑/↓ navigate · ←/→ week · q quit"

	case data.HoursErrMsg:
		m.statusLine = "hours error: " + msg.Err.Error()
		m.loading = false

	case data.LecturesFetchedMsg:
		m.lectures = msg.Lectures
		m.rebuildGrid()
		m.loading = false
		m.statusLine = "↑/↓ navigate · ←/→ week · q quit"

	case data.LecturesErrMsg:
		m.statusLine = "lectures error: " + msg.Err.Error()
		m.loading = false

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursorRow > 0 {
				m.cursorRow--
			}
		case "down", "j":
			max := len(m.grid.visibleIdx) - 1
			if max < 0 {
				max = 0
			}
			if m.cursorRow < max {
				m.cursorRow++
			}
		case "left", "h":
			m.weekOffset--
			m.weekStart = currentWeekStart().AddDate(0, 0, m.weekOffset*7)
			m.loading = true
			m.statusLine = "loading..."
			return m, tea.Batch(data.FetchHours(), data.FetchLectures())
		case "right", "l":
			m.weekOffset++
			m.weekStart = currentWeekStart().AddDate(0, 0, m.weekOffset*7)
			m.loading = true
			m.statusLine = "loading..."
			return m, tea.Batch(data.FetchHours(), data.FetchLectures())
		}
	}

	return m, nil
}

func (m *model) rebuildGrid() {
	weekEnd := m.weekStart.AddDate(0, 0, 7)
	var weekLectures []data.Lecture
	for _, lec := range m.lectures {
		if len(lec.Date) < 10 {
			continue
		}
		d, err := time.Parse("2006-01-02", lec.Date)
		if err != nil {
			continue
		}
		if !d.Before(m.weekStart) && d.Before(weekEnd) {
			weekLectures = append(weekLectures, lec)
		}
	}
	m.grid = buildGrid(m.hours, weekLectures)
	if m.cursorRow >= len(m.grid.visibleIdx) {
		m.cursorRow = len(m.grid.visibleIdx) - 1
	}
	if m.cursorRow < 0 {
		m.cursorRow = 0
	}
}

func (m model) View() string {
	w := m.width
	if w == 0 {
		w = 80
	}

	weekEnd := m.weekStart.AddDate(0, 0, 6)
	weekLabel := fmt.Sprintf("Week: %s - %s",
		m.weekStart.Format("Jan 02"),
		weekEnd.Format("Jan 02, 2006"))

	title := titleStyle.Render("caly")
	week := weekStyle.Render(weekLabel)
	gap := w - lipgloss.Width(title) - lipgloss.Width(week)
	gap = min(gap, 1)
	header := title + strings.Repeat(" ", gap) + week

	gridStr := ""
	if m.loading {
		gridStr = lipgloss.NewStyle().
			Foreground(lipgloss.Color("241")).
			Padding(1, 2).
			Render("  Loading...")
	} else {
		gridStr = lipgloss.NewStyle().Padding(1, 1).Render(
			renderGrid(m.grid, m.cursorRow, m.weekStart, w),
		)
	}

	detail := ""
	if !m.loading {
		lec := lectureAtCursor(m.grid, m.cursorRow)
		if lec != nil {
			color := ColorForSubject(lec.SubjectCode)
			colorBlock := lipgloss.NewStyle().
				Background(color).
				Foreground(lipgloss.Color("#1a1a1a")).
				Bold(true).
				Padding(0, 1).
				Render(lec.SubjectCode)

			room := lec.Room
			lecturers := strings.Join(lec.Lecturers, ", ")

			detail = detailStyle.Render(
				colorBlock + "  " +
					detailLabelStyle.Render(lec.SubjectName) +
					"  ·  " + room +
					"  ·  " + lecturers,
			)
		}
	}

	status := statusBarStyle.Render(m.statusLine)

	return lipgloss.JoinVertical(lipgloss.Left,
		header,
		gridStr,
		detail,
		status,
	)
}

func Run() error {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	_, err := p.Run()
	return err
}
