package cmd

import (
	"fmt"
	"slices"
	"time"

	"github.com/mmmcclimon/toggl-go/internal/toggl"
	"github.com/spf13/cobra"
)

type TimelineCommand struct{}

func (cmd *TimelineCommand) Cobra() *cobra.Command {
	return &cobra.Command{
		Use:     "timeline [DATE]",
		Aliases: []string{"tl"},
		Short:   "wait what did I do again?",
	}
}

func (cmd *TimelineCommand) Run(tc *toggl.Client, args []string) error {
	var start time.Time
	switch len(args) {
	case 0:
		start = startOfToday()

	case 1:
		var err error
		start, err = time.ParseInLocation(time.DateOnly, args[0], time.Local)
		if err != nil {
			return fmt.Errorf("could not parse %q as a date: %w", args[0], err)
		}

	default:
		return fmt.Errorf("need only one date to check")
	}

	entries, err := tc.TimeEntries(start, start.Add(24*time.Hour))
	if err != nil {
		return err
	}

	if len(entries) == 0 {
		fmt.Println("Nothing logged today.")
		return nil
	}

	return cmd.printLog(entries)
}

func (cmd *TimelineCommand) printLog(entries []*toggl.Timer) error {
	slices.Reverse(entries)

	var total time.Duration

	for _, t := range entries {
		end := t.End.Local().Format("15:04")
		if t.End.IsZero() {
			end = ""
		}

		fmt.Printf("%s–%5s  %s\n",
			t.Start.Local().Format("15:04"),
			end,
			t.OnelineDesc(),
		)

		total += t.Duration()
	}

	fmt.Println("-----------")
	fmt.Printf("%0.2fh total (%s)\n", total.Hours(), total)
	return nil
}
