package cmd

import (
	"fmt"
	"time"

	"github.com/mmmcclimon/toggl-go/internal/toggl"
	"github.com/spf13/cobra"
)

type WeekCommand struct{}

func (cmd WeekCommand) Cobra() *cobra.Command {
	return &cobra.Command{
		Use:   "week",
		Short: "how's the week been?",
	}
}

func (cmd WeekCommand) Run(tc *toggl.Client, args []string) error {
	end := time.Now()
	start := startOfToday()

	// Back up til we hit a Monday (the week starts on Monday, come at me.)
	for start.Local().Weekday() != time.Monday {
		start = start.Add(-24 * time.Hour)
	}

	entries, err := tc.TimeEntries(start, end)
	if err != nil {
		return err
	}

	if len(entries) == 0 {
		fmt.Println("Nothing logged this week.")
		return nil
	}

	toggl.PrintEntryList(entries)
	return nil
}
