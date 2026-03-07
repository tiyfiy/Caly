package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/tiyfiy/caly/internal/data"
	"github.com/tiyfiy/caly/internal/ui"
)

func main() {
	discover := flag.Bool("discover", false, "discover your iCloud calendar paths")
	flag.Parse()

	if *discover {
		username := os.Getenv("ICLOUD_USERNAME")
		password := os.Getenv("ICLOUD_APP_PASSWORD")

		if username == "" || password == "" {
			fmt.Fprintln(os.Stderr, "error: ICLOUD_USERNAME and ICLOUD_APP_PASSWORD environment variables must be set")
			os.Exit(1)
		}

		paths, err := data.DiscoverCalendarPath(username, password)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("Your iCloud calendars:")
		fmt.Println()
		for _, p := range paths {
			fmt.Println(" ", p)
		}
		fmt.Println()
		fmt.Println("Set the path part (after the →) as ICLOUD_CALENDAR_PATH")
		return
	}

	if err := ui.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
