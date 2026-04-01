package ui

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/tiyfiy/caly/internal/data"
)

type gridCell struct {
	subjectCode string
	isStart     bool
	lecture     *data.Lecture
}

type gridData struct {
	cells      map[[2]int][]gridCell
	visibleIdx []int
	hours      []data.Hour
}

func buildGrid(hours []data.Hour, lectures []data.Lecture) gridData {
	cells := make(map[[2]int][]gridCell)
	usedRows := make(map[int]bool)

	for i := range lectures {
		lec := &lectures[i]
		if len(lec.Start) < 16 || len(lec.End) < 16 || len(lec.Date) < 10 {
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

		lecStart := lec.Start[11:16]
		lecEnd := lec.End[11:16]

		first := true
		for i, h := range hours {
			slotStart := h.Start[:5]
			slotEnd := h.End[:5]

			if slotStart >= lecStart && slotEnd <= lecEnd {
				key := [2]int{i, col}
				cells[key] = append(cells[key], gridCell{
					subjectCode: lec.SubjectCode,
					isStart:     first,
					lecture:     lec,
				})
				usedRows[i] = true
				first = false
			}
		}
	}

	visible := make(map[int]bool)
	for r := range usedRows {
		visible[r] = true
		if r > 0 {
			visible[r-1] = true
		}
		if r < len(hours)-1 {
			visible[r+1] = true
		}
	}

	var visibleIdx []int
	for i := range hours {
		if visible[i] {
			visibleIdx = append(visibleIdx, i)
		}
	}

	return gridData{cells: cells, visibleIdx: visibleIdx, hours: hours}
}

func currentTimeRow(hours []data.Hour) int {
	now := time.Now()
	nowStr := fmt.Sprintf("%02d:%02d", now.Hour(), now.Minute())
	for i, h := range hours {
		if nowStr >= h.Start[:5] && nowStr < h.End[:5] {
			return i
		}
	}
	return -1
}

func renderGrid(g gridData, cursorRow int, weekStart time.Time, width int) string {
	if len(g.visibleIdx) == 0 {
		return lipgloss.NewStyle().
			Foreground(lipgloss.Color("#4B5563")).
			Padding(1, 0).
			Render("No lectures scheduled this week")
	}

	hasSat, hasSun := false, false
	for key := range g.cells {
		if key[1] == 6 {
			hasSat = true
		}
		if key[1] == 7 {
			hasSun = true
		}
	}
	dayCols := 5
	if hasSat {
		dayCols = 6
	}
	if hasSun {
		dayCols = 7
	}

	const prefixW = 2
	const timeColW = 7
	const sepW = 1

	available := width - prefixW - timeColW - dayCols*sepW
	available = max(dayCols*10, available)

	cellW := available / dayCols

	// Detect today's column (1-indexed, Mon=1)
	now := time.Now()
	todayCol := -1
	for d := 0; d < dayCols; d++ {
		day := weekStart.AddDate(0, 0, d)
		if day.Year() == now.Year() && day.YearDay() == now.YearDay() {
			todayCol = d + 1
			break
		}
	}

	nowRow := currentTimeRow(g.hours)

	// ── Styles ────────────────────────────────────────────────────────────────

	headerBg := lipgloss.Color("#0F172A")    // slate-900
	todayBg := lipgloss.Color("#1E1B4B")     // indigo-950
	todayFg := lipgloss.Color("#A5B4FC")     // indigo-300
	headerFg := lipgloss.Color("#D1D5DB")    // gray-300
	mutedFg := lipgloss.Color("#6B7280")     // gray-500
	currentFg := lipgloss.Color("#F59E0B")   // amber-400
	accentFg := lipgloss.Color("#818CF8")    // indigo-400
	brightFg := lipgloss.Color("#F9FAFB")    // gray-50
	separatorFg := lipgloss.Color("#1F2937") // gray-800
	cursorRowBg := lipgloss.Color("#1E293B") // slate-800
	lecCursorBg := lipgloss.Color("#E0E7FF") // indigo-100
	lecCursorFg := lipgloss.Color("#1E1B4B") // indigo-950
	textDark := lipgloss.Color("#111827")    // gray-900

	dayHeaderStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(headerFg).
		Background(headerBg).
		Width(cellW).
		Align(lipgloss.Center)

	todayHeaderStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(todayFg).
		Background(todayBg).
		Width(cellW).
		Align(lipgloss.Center)

	timeHeaderBg := lipgloss.NewStyle().
		Background(headerBg).
		Width(prefixW + timeColW)

	timeCellStyle := lipgloss.NewStyle().
		Foreground(mutedFg).
		Width(timeColW).
		Align(lipgloss.Right).
		Padding(0, 1)

	sep := lipgloss.NewStyle().
		Foreground(separatorFg).
		Render("│")

	headerSep := lipgloss.NewStyle().
		Background(headerBg).
		Foreground(separatorFg).
		Render("│")

	// ── Header row ────────────────────────────────────────────────────────────

	dayNames := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	var headerParts []string
	headerParts = append(headerParts, timeHeaderBg.Render(""))

	for d := 0; d < dayCols; d++ {
		day := weekStart.AddDate(0, 0, d)
		label := fmt.Sprintf("%s %02d", dayNames[d], day.Day())
		col := d + 1

		if col == todayCol {
			headerParts = append(headerParts, headerSep+todayHeaderStyle.Render(label))
		} else {
			headerParts = append(headerParts, headerSep+dayHeaderStyle.Render(label))
		}
	}

	header := strings.Join(headerParts, "")

	sepLine := lipgloss.NewStyle().
		Foreground(separatorFg).
		Render(strings.Repeat("─", width))

	// ── Time rows ─────────────────────────────────────────────────────────────

	var rows []string
	rows = append(rows, header)
	rows = append(rows, sepLine)

	for vi, hourIdx := range g.visibleIdx {
		h := g.hours[hourIdx]
		timeLabel := h.Start[:5]

		isCurrentTime := hourIdx == nowRow
		isCursor := vi == cursorRow

		tStyle := timeCellStyle
		prefix := "  "

		switch {
		case isCursor:
			prefix = lipgloss.NewStyle().Foreground(accentFg).Bold(true).Render("▶ ")
			tStyle = tStyle.Bold(true).Foreground(brightFg)
		case isCurrentTime:
			prefix = lipgloss.NewStyle().Foreground(currentFg).Render("◆ ")
			tStyle = tStyle.Foreground(currentFg)
		}

		var rowParts []string
		rowParts = append(rowParts, prefix+tStyle.Render(timeLabel))

		for col := 1; col <= dayCols; col++ {
			key := [2]int{hourIdx, col}
			subCells, hasLecture := g.cells[key]

			// Last column gets the remainder width to fill exactly
			colW := cellW
			if col == dayCols {
				colW = available - cellW*(dayCols-1)
			}

			var cell string
			if hasLecture {
				n := len(subCells)
				subW := colW / n

				var parts []string
				for si, sc := range subCells {
					w := subW
					if si == n-1 {
						w = colW - subW*(n-1)
					}

					bg := ColorForSubject(sc.subjectCode)
					cellStyle := lipgloss.NewStyle().
						Width(w).
						Align(lipgloss.Center).
						Background(bg).
						Foreground(textDark).
						Bold(sc.isStart)

					if isCursor {
						cellStyle = cellStyle.
							Background(lecCursorBg).
							Foreground(lecCursorFg).
							Bold(true)
					}

					text := ""
					if sc.isStart {
						text = sc.subjectCode
					}
					parts = append(parts, cellStyle.Render(text))
				}
				cell = lipgloss.JoinHorizontal(lipgloss.Top, parts...)
			} else {
				cs := lipgloss.NewStyle().Width(colW)
				if isCursor {
					cs = cs.Background(cursorRowBg)
				}
				cell = cs.Render("")
			}

			rowParts = append(rowParts, sep+cell)
		}

		rows = append(rows, strings.Join(rowParts, ""))
	}

	return lipgloss.JoinVertical(lipgloss.Left, rows...)
}

func lectureAtCursor(g gridData, cursorRow int) *data.Lecture {
	if cursorRow < 0 || cursorRow >= len(g.visibleIdx) {
		return nil
	}
	hourIdx := g.visibleIdx[cursorRow]
	for col := 1; col <= 7; col++ {
		if subCells, ok := g.cells[[2]int{hourIdx, col}]; ok && len(subCells) > 0 {
			return subCells[0].lecture
		}
	}
	return nil
}

func lectureAtCursorCol(g gridData, cursorRow, cursorCol int) *data.Lecture {
	if cursorRow < 0 || cursorRow >= len(g.visibleIdx) {
		return nil
	}
	hourIdx := g.visibleIdx[cursorRow]
	if subCells, ok := g.cells[[2]int{hourIdx, cursorCol}]; ok && len(subCells) > 0 {
		return subCells[0].lecture
	}
	return nil
}
