package ui

import (
	"fmt"
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
		return lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Render("  No lectures this week")
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

	timeColW := 13
	availableW := width - timeColW - 2
	availableW = max(availableW, dayCols*8)
	cellW := availableW / dayCols

	nowRow := currentTimeRow(g.hours)

	headerStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("255")).
		Background(lipgloss.Color("236")).
		Width(cellW).
		Align(lipgloss.Center).
		Padding(0, 1)

	timeHeaderStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("255")).
		Background(lipgloss.Color("236")).
		Width(timeColW).
		Align(lipgloss.Center)

	timeCellStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("245")).
		Width(timeColW).
		Align(lipgloss.Right).
		Padding(0, 1)

	emptyCellStyle := lipgloss.NewStyle().
		Width(cellW).
		Align(lipgloss.Center)

	cursorMarker := lipgloss.NewStyle().
		Foreground(lipgloss.Color("226"))

	cursorRowBg := lipgloss.Color("235")
	_ = cursorRowBg

	dayNames := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	var headerCells []string
	headerCells = append(headerCells, timeHeaderStyle.Render(""))
	for d := 0; d < dayCols; d++ {
		day := weekStart.AddDate(0, 0, d)
		label := fmt.Sprintf("%s %02d", dayNames[d], day.Day())
		headerCells = append(headerCells, headerStyle.Render(label))
	}
	header := lipgloss.JoinHorizontal(lipgloss.Top, headerCells...)

	var rows []string
	rows = append(rows, header)

	for vi, hourIdx := range g.visibleIdx {
		h := g.hours[hourIdx]
		timeLabel := fmt.Sprintf("%s-%s", h.Start[:5], h.End[:5])

		isCurrentTime := hourIdx == nowRow
		isCursor := vi == cursorRow

		tStyle := timeCellStyle
		if isCurrentTime {
			tStyle = tStyle.Foreground(lipgloss.Color("226"))
		}

		prefix := "  "
		if isCursor {
			prefix = cursorMarker.Render("→ ")
			tStyle = tStyle.Bold(true).Foreground(lipgloss.Color("255"))
		} else if isCurrentTime {
			prefix = cursorMarker.Render("• ")
		}

		var rowCells []string
		rowCells = append(rowCells, prefix+tStyle.Render(timeLabel))

		for col := 1; col <= dayCols; col++ {
			key := [2]int{hourIdx, col}
			subCells, hasLecture := g.cells[key]

			if hasLecture {
				n := len(subCells)
				subW := cellW / n

				var parts []string
				for si, cell := range subCells {
					w := subW
					if si == n-1 {
						w = cellW - subW*(n-1)
					}

					bg := ColorForSubject(cell.subjectCode)
					style := lipgloss.NewStyle().
						Width(w).
						Align(lipgloss.Center).
						Background(bg).
						Foreground(lipgloss.Color("#1a1a1a")).
						Bold(cell.isStart)

					if isCursor {
						style = style.
							Background(lipgloss.Color("#ffffff")).
							Foreground(lipgloss.Color("#1a1a1a")).
							Bold(true)
					}

					text := ""
					if cell.isStart {
						text = cell.subjectCode
					}
					parts = append(parts, style.Render(text))
				}
				rowCells = append(rowCells, lipgloss.JoinHorizontal(lipgloss.Top, parts...))
			} else {
				style := emptyCellStyle
				if isCursor {
					style = style.Foreground(lipgloss.Color("238"))
				}
				rowCells = append(rowCells, style.Render("·"))
			}
		}
		rows = append(rows, lipgloss.JoinHorizontal(lipgloss.Top, rowCells...))
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
