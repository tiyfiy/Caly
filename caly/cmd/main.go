package main

import (
	"fmt"
	"os"

	"github.com/tiyfiy/caly/internal/ui"
)

func main() {
	if err := ui.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
