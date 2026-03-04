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
		{Title: "Mon", Width: 14},
		{Title: "Tue", Width: 14},
		{Title: "Wed", Width: 14},
		{Title: "Thu", Width: 14},
		{Title: "Fri", Width: 14},
		{Title: "Sat", Width: 14},
		{Title: "Sun", Width: 14},
	}

	// cellKey identifies a single cell by row and column index
	type cellKey struct{ row, col int }
	grid := make(map[cellKey]string)

	for _, lec := range lectures {
		if len(lec.Start) < 16 || len(lec.End) < 16 || len(lec.Date) < 10 {
			continue
		}

		t, err := time.Parse("2006-01-02", lec.Date)
		if err != nil {
			continue
		}

		// time.Weekday: Sun=0 Mon=1 ... Sat=6
		// table cols:   Mon=1 Tue=2 ... Sat=6 Sun=7
		wd := int(t.Weekday())
		col := wd
		if wd == 0 {
			col = 7
		}

		// slice timezone-aware ISO string: "2026-03-05T15:15:00+01:00" -> "15:15"
		lecStart := lec.Start[11:16]
		lecEnd := lec.End[11:16]

		first := true
		for i, h := range hours {
			slotStart := h.Start[:5] // "15:15" from "15:15:00"
			slotEnd := h.End[:5]

			// slot is covered when it falls fully within the lecture window
			if slotStart >= lecStart && slotEnd <= lecEnd {
				if _, exists := grid[cellKey{i, col}]; exists {
					// another lecture already claimed this cell, skip
					continue
				}
				if first {
					// first slot: show subject code with a start marker
					grid[cellKey{i, col}] = fmt.Sprintf("▶ %s", lec.SubjectCode)
					first = false
				} else {
					// continuation slot: vertical bar to show the class is still going
					grid[cellKey{i, col}] = "  │"
				}
			}
		}
	}

	rows := make([]table.Row, len(hours))
	for i, h := range hours {
		label := fmt.Sprintf("%s-%s", h.Start[:5], h.End[:5])
		row := table.Row{label, "", "", "", "", "", "", ""}
		for col := 1; col <= 7; col++ {
			if cell, ok := grid[cellKey{i, col}]; ok {
				row[col] = cell
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
