package data

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

const iCloudCalDAVBase = "https://caldav.icloud.com"

type CalendarPushedMsg struct{ Count int }
type CalendarErrMsg struct{ Err error }

type CalDAVConfig struct {
	Username     string
	AppPassword  string
	CalendarPath string
	CalDAVHost   string
}

type davHref struct {
	Value string `xml:",chardata"`
}
type davPropStat struct {
	Prop davProp `xml:"prop"`
}
type davProp struct {
	CurrentUserPrincipal struct {
		Href davHref `xml:"href"`
	} `xml:"current-user-principal"`
	CalendarHomeSet struct {
		Href davHref `xml:"href"`
	} `xml:"calendar-home-set"`
	DisplayName  string `xml:"displayname"`
	ResourceType struct {
		Calendar *struct{} `xml:"calendar"`
	} `xml:"resourcetype"`
}
type davMultiStatus struct {
	Responses []struct {
		Href     string      `xml:"href"`
		PropStat davPropStat `xml:"propstat"`
	} `xml:"response"`
}

var httpClient = &http.Client{Timeout: 10 * time.Second}

func propfind(url, username, password, depth, body string) (*davMultiStatus, error) {
	req, err := http.NewRequest("PROPFIND", url, strings.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(username, password)
	req.Header.Set("Depth", depth)
	req.Header.Set("Content-Type", "application/xml")

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		return nil, fmt.Errorf("unauthorized — check your username and app-specific password")
	}
	if resp.StatusCode != 207 {
		return nil, fmt.Errorf("unexpected status %d", resp.StatusCode)
	}

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var ms davMultiStatus
	if err := xml.Unmarshal(raw, &ms); err != nil {
		return nil, err
	}
	return &ms, nil
}

func PushToCalendar(lectures []Lecture, cfg CalDAVConfig) tea.Cmd {
	return func() tea.Msg {
		count := 0
		for _, lec := range lectures {
			if err := pushEventCalDAV(lec, cfg); err != nil {
				return CalendarErrMsg{Err: fmt.Errorf("failed to push %s: %w", lec.SubjectCode, err)}
			}
			count++
		}
		return CalendarPushedMsg{Count: count}
	}
}

func pushEventCalDAV(lec Lecture, cfg CalDAVConfig) error {
	uid := fmt.Sprintf("%s-%s@caly", lec.SubjectCode, lec.Date)

	start, err := formatForICS(lec.Start)
	if err != nil {
		return fmt.Errorf("invalid start: %w", err)
	}
	end, err := formatForICS(lec.End)
	if err != nil {
		return fmt.Errorf("invalid end: %w", err)
	}

	lecturers := strings.Join(lec.Lecturers, ", ")
	if lecturers == "" {
		lecturers = "—"
	}

	ics := fmt.Sprintf(
		"BEGIN:VCALENDAR\r\nVERSION:2.0\r\nPRODID:-//caly//EN\r\n"+
			"BEGIN:VEVENT\r\nUID:%s\r\nDTSTART:%s\r\nDTEND:%s\r\n"+
			"SUMMARY:%s\r\nLOCATION:%s\r\nDESCRIPTION:%s - %s\r\n"+
			"END:VEVENT\r\nEND:VCALENDAR\r\n",
		uid, start, end, lec.SubjectName, lec.Room, lec.SubjectCode, lecturers,
	)

	host := cfg.CalDAVHost
	if host == "" {
		host = iCloudCalDAVBase
	}
	calPath := cfg.CalendarPath
	if calPath != "" && !strings.HasSuffix(calPath, "/") {
		calPath += "/"
	}
	safeUID := strings.ReplaceAll(uid, "@", "%40")
	url := host + calPath + safeUID + ".ics"

	req, err := http.NewRequest(http.MethodPut, url, strings.NewReader(ics))
	if err != nil {
		return err
	}
	req.SetBasicAuth(cfg.Username, cfg.AppPassword)
	req.Header.Set("Content-Type", "text/calendar; charset=utf-8")
	req.Header.Set("If-None-Match", "*")

	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusPreconditionFailed {
		return nil
	}
	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("unexpected status %d for event %s", resp.StatusCode, uid)
	}
	return nil
}

func formatForICS(iso string) (string, error) {
	if len(iso) < 19 {
		return "", fmt.Errorf("timestamp too short: %q", iso)
	}
	date := strings.ReplaceAll(iso[:10], "-", "")
	clock := strings.ReplaceAll(iso[11:19], ":", "")
	return date + "T" + clock, nil
}

func DiscoverCalendarPath(username, appPassword string) ([]string, error) {
	// 1. find principal
	ms, err := propfind(iCloudCalDAVBase+"/", username, appPassword, "0",
		`<?xml version="1.0"?><D:propfind xmlns:D="DAV:"><D:prop><D:current-user-principal/></D:prop></D:propfind>`)
	if err != nil {
		return nil, fmt.Errorf("discover principal: %w", err)
	}

	var principalURL string
	for _, r := range ms.Responses {
		if h := r.PropStat.Prop.CurrentUserPrincipal.Href.Value; h != "" {
			principalURL = h
			break
		}
	}
	if principalURL == "" {
		return nil, fmt.Errorf("no principal found")
	}

	// 2. find calendar home set
	ms, err = propfind(iCloudCalDAVBase+principalURL, username, appPassword, "0",
		`<?xml version="1.0"?><D:propfind xmlns:D="DAV:" xmlns:C="urn:ietf:params:xml:ns:caldav"><D:prop><C:calendar-home-set/></D:prop></D:propfind>`)
	if err != nil {
		return nil, fmt.Errorf("discover home set: %w", err)
	}

	var homeSet string
	for _, r := range ms.Responses {
		if h := r.PropStat.Prop.CalendarHomeSet.Href.Value; h != "" {
			homeSet = h
			break
		}
	}
	if homeSet == "" {
		return nil, fmt.Errorf("no calendar home set found")
	}

	homeSetURL := homeSet
	if !strings.HasPrefix(homeSet, "http") {
		homeSetURL = iCloudCalDAVBase + homeSet
	}

	// 3. list calendars
	ms, err = propfind(homeSetURL, username, appPassword, "1",
		`<?xml version="1.0"?><D:propfind xmlns:D="DAV:"><D:prop><D:displayname/><D:resourcetype/></D:prop></D:propfind>`)
	if err != nil {
		return nil, fmt.Errorf("list calendars: %w", err)
	}

	var paths []string
	for _, r := range ms.Responses {
		if r.PropStat.Prop.ResourceType.Calendar != nil {
			paths = append(paths, fmt.Sprintf("%s  →  %s", r.PropStat.Prop.DisplayName, r.Href))
		}
	}

	if len(paths) == 0 {
		return nil, fmt.Errorf("no calendars found")
	}
	return paths, nil
}
