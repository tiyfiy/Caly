package ui

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
	"github.com/tiyfiy/caly/internal/data"
)

const reservedLines = 6

func newTable(hours []data.Hour, lectures []data.Lecture, height int) table.Model {
	columns := []table.Column{
		{Title: "Time", Width: 11},
		{Title: "Mon", Width: 12},
		{Title: "Tue", Width: 12},
		{Title: "Wed", Width: 12},
		{Title: "Thu", Width: 12},
		{Title: "Fri", Width: 12},
		{Title: "Sat", Width: 12},
		{Title: "Sun", Width: 12},
	}

	type cellKey struct{ row, col int }
	grid := make(map[cellKey]string)

	for _, lec := range lectures {
		if len(lec.Start) < 16 || len(lec.Date) < 10 {
			continue
		}
		t, err := time.Parse("2006-01-02", lec.Date)
		if err != nil {
			continue
		}
		wd := int(t.Weekday())
		col := wd
		if wd == 0 {
			col = 7
		}

		lectureStart := lec.Start[11:16]
		for i, h := range hours {
			if h.Start[:5] == lectureStart {
				grid[cellKey{i, col}] = lec.SubjectCode
				break
			}
		}
	}

	rows := make([]table.Row, len(hours))
	for i, h := range hours {
		label := fmt.Sprintf("%s-%s", h.Start[:5], h.End[:5])
		row := table.Row{label, "", "", "", "", "", "", ""}
		for col := 1; col <= 7; col++ {
			if subj, ok := grid[cellKey{i, col}]; ok {
				row[col] = subj
			}
		}
		rows[i] = row
	}

	tableHeight := height - reservedLines
	if tableHeight < 1 {
		tableHeight = 1
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(tableHeight),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("62")).
		BorderBottom(true).
		BorderLeft(true).
		Bold(true).
		Foreground(lipgloss.Color("62"))
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(true)
	s.Cell = s.Cell.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderLeft(true).
		BorderBottom(true).
		Foreground(lipgloss.Color("252"))
	t.SetStyles(s)

	return t
}
