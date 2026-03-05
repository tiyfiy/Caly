package ui

import "github.com/charmbracelet/lipgloss"

var subjectPalette = []lipgloss.Color{
	lipgloss.Color("#87CEEB"), // sky blue
	lipgloss.Color("#98FB98"), // pale green
	lipgloss.Color("#DDA0DD"), // plum
	lipgloss.Color("#F0E68C"), // khaki
	lipgloss.Color("#FFA07A"), // light salmon
	lipgloss.Color("#B0C4DE"), // light steel blue
	lipgloss.Color("#FFB6C1"), // light pink
	lipgloss.Color("#20B2AA"), // light sea green
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
