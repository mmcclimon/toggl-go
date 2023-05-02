package cmd

import (
	"fmt"
	"time"

	t "github.com/mmmcclimon/toggl-go/internal/toggl"
	"github.com/spf13/cobra"
)

type TodayCommand struct{}

func (cmd TodayCommand) Cobra() *cobra.Command {
	return &cobra.Command{
		Use:   "today",
		Short: "what are the things you've done today?",
	}
}

func (cmd TodayCommand) Run(toggl *t.Toggl, args []string) error {
	start := startOfToday()
	end := time.Now()

	entries, err := toggl.TimeEntries(start, end)
	if err != nil {
		return err
	}

	if len(entries) == 0 {
		fmt.Println("Nothing logged today.")
		return nil
	}

	t.PrintEntryList(entries)
	return nil
}
