package ui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
	"github.com/tiyfiy/caly/internal/data"
)

func newTable(hours []data.Hour, height int) table.Model {
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

	rows := make([]table.Row, len(hours))
	for i, h := range hours {
		start := h.Start[:5]
		end := h.End[:5]
		label := fmt.Sprintf("%s-%s", start, end)
		rows[i] = table.Row{label, "", "", "", "", "", "", ""}
	}

	tableHeight := height - 6
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
