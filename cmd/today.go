package cmd

import (
	"fmt"
	"time"

	"github.com/mmmcclimon/toggl-go/internal/toggl"
	"github.com/spf13/cobra"
)

type TodayCommand struct{}

func (cmd *TodayCommand) Cobra() *cobra.Command {
	return &cobra.Command{
		Use:   "today",
		Short: "what are the things you've done today?",
	}
}

func (cmd *TodayCommand) Run(tc *toggl.Client, args []string) error {
	start := startOfToday()
	end := time.Now()

	entries, err := tc.TimeEntries(start, end)
	if err != nil {
		return err
	}

	if len(entries) == 0 {
		fmt.Println("Nothing logged today.")
		return nil
	}

	toggl.PrintEntryList(entries)
	return nil
}
