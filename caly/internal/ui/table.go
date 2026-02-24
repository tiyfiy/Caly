package ui

import (
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

var timeSlots = []string{
	"08:00", "09:00", "10:00", "11:00", "12:00",
	"13:00", "14:00", "15:00", "16:00", "17:00",
	"18:00", "19:00", "20:00",
}

func newTable() table.Model {
	columns := []table.Column{
		{Title: "Time", Width: 6},
		{Title: "Mon", Width: 12},
		{Title: "Tue", Width: 12},
		{Title: "Wed", Width: 12},
		{Title: "Thu", Width: 12},
		{Title: "Fri", Width: 12},
		{Title: "Sat", Width: 12},
		{Title: "Sun", Width: 12},
	}

	rows := make([]table.Row, len(timeSlots))
	for i, slot := range timeSlots {
		rows[i] = table.Row{slot, "", "", "", "", "", "", ""}
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(len(timeSlots)),
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
