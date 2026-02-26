package data

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
)

type HoursFetchedMsg struct {
	Hours []Hour
}

type HoursErrMsg struct {
	Err error
}

func FetchHours() tea.Cmd {
	return func() tea.Msg {
		cmd := exec.Command("scraper", "--hours")
		if err := cmd.Run(); err != nil {
			return HoursErrMsg{Err: fmt.Errorf("scraper failed: %w", err)}
		}

		bytes, err := os.ReadFile("hours.json")
		if err != nil {
			return HoursErrMsg{Err: fmt.Errorf("could not read hours.json: %w", err)}
		}

		var hours []Hour
		if err := json.Unmarshal(bytes, &hours); err != nil {
			return HoursErrMsg{Err: fmt.Errorf("could not parse hours.json: %w", err)}
		}

		return HoursFetchedMsg{Hours: hours}
	}
}
