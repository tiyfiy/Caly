package ui

import (
	"fmt"
	"strings"
	"time"

	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/tiyfiy/caly/internal/data"
)

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#818CF8")).
			Padding(0, 1)

	weekStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#9CA3AF")).
			Padding(0, 1)

	dividerStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#374151"))

	detailStyle = lipgloss.NewStyle().
			Padding(0, 2)

	detailBadgeStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("#111827")).
				Padding(0, 1)

	detailNameStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#F9FAFB"))

	detailMetaStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#6B7280"))

	statusBarStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#4B5563")).
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
		m.statusLine = "↑↓ navigate   ←→ week   q quit"

	case data.HoursErrMsg:
		m.statusLine = "hours error: " + msg.Err.Error()
		m.loading = false

	case data.LecturesFetchedMsg:
		m.lectures = msg.Lectures
		m.rebuildGrid()
		m.loading = false
		m.statusLine = "↑↓ navigate   ←→ week   p push   q quit"

	case data.LecturesErrMsg:
		m.statusLine = "lectures error: " + msg.Err.Error()
		m.loading = false

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case data.CalendarPushedMsg:
		m.statusLine = fmt.Sprintf("✓ pushed %d events to Calendar", msg.Count)

	case data.CalendarErrMsg:
		m.statusLine = "calendar error: " + msg.Err.Error()

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
			from := m.weekStart.Format("2006-01-02")
			to := m.weekStart.AddDate(0, 0, 6).Format("2006-01-02")
			return m, tea.Batch(data.FetchHours(), data.FetchLecturesWithDate(from, to))
		case "right", "l":
			m.weekOffset++
			m.weekStart = currentWeekStart().AddDate(0, 0, m.weekOffset*7)
			m.loading = true
			m.statusLine = "loading..."
			from := m.weekStart.Format("2006-01-02")
			to := m.weekStart.AddDate(0, 0, 6).Format("2006-01-02")
			return m, tea.Batch(data.FetchHours(), data.FetchLecturesWithDate(from, to))
		case "p":
			cfg := data.CalDAVConfig{
				Username:     os.Getenv("ICLOUD_USERNAME"),
				AppPassword:  os.Getenv("ICLOUD_APP_PASSWORD"),
				CalendarPath: os.Getenv("ICLOUD_CALENDAR_PATH"),
				CalDAVHost:   os.Getenv("ICLOUD_CALDAV_HOST"),
			}
			return m, data.PushToCalendar(m.lectures, cfg)
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

	// ── Header ────────────────────────────────────────────────────────────────

	title := titleStyle.Render("caly")

	weekEnd := m.weekStart.AddDate(0, 0, 6)
	weekLabel := fmt.Sprintf("%s – %s",
		m.weekStart.Format("Jan 02"),
		weekEnd.Format("Jan 02, 2006"))
	week := weekStyle.Render(weekLabel)

	gap := w - lipgloss.Width(title) - lipgloss.Width(week)
	gap = max(1, gap)

	header := title + strings.Repeat(" ", gap) + week

	divider := dividerStyle.Render(strings.Repeat("─", w))

	// ── Grid ──────────────────────────────────────────────────────────────────

	gridW := w - 2
	gridW = max(20, gridW)

	var gridStr string
	if m.loading {
		gridStr = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#6B7280")).
			Padding(2, 2).
			Render("loading...")
	} else {
		gridStr = lipgloss.NewStyle().
			Padding(1, 1).
			Render(renderGrid(m.grid, m.cursorRow, m.weekStart, gridW))
	}

	// ── Detail ────────────────────────────────────────────────────────────────

	detail := ""
	if !m.loading {
		lec := lectureAtCursor(m.grid, m.cursorRow)
		if lec != nil {
			color := ColorForSubject(lec.SubjectCode)
			badge := detailBadgeStyle.Background(color).Render(lec.SubjectCode)
			name := detailNameStyle.Render(lec.SubjectName)
			meta := detailMetaStyle.Render(
				"   ·  " + lec.Room + "   ·  " + strings.Join(lec.Lecturers, ", "),
			)
			detail = detailStyle.Render(badge + "  " + name + meta)
		}
	}

	// ── Status bar ────────────────────────────────────────────────────────────

	status := statusBarStyle.Render(m.statusLine)

	return lipgloss.JoinVertical(lipgloss.Left,
		header,
		divider,
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
