package ui

import "github.com/charmbracelet/lipgloss"

var subjectPalette = []lipgloss.Color{
	lipgloss.Color("#93C5FD"), // blue-300
	lipgloss.Color("#86EFAC"), // green-300
	lipgloss.Color("#FDE68A"), // amber-200
	lipgloss.Color("#FCA5A5"), // red-300
	lipgloss.Color("#C4B5FD"), // violet-300
	lipgloss.Color("#5EEAD4"), // teal-300
	lipgloss.Color("#F9A8D4"), // pink-300
	lipgloss.Color("#FED7AA"), // orange-200
}

var subjectColorMap = map[string]lipgloss.Color{}

func ColorForSubject(code string) lipgloss.Color {
	if c, ok := subjectColorMap[code]; ok {
		return c
	}
	c := subjectPalette[len(subjectColorMap)%len(subjectPalette)]
	subjectColorMap[code] = c
	return c
}
