package cmd

import (
	"fmt"
	"slices"
	"time"

	"github.com/mmmcclimon/toggl-go/internal/toggl"
	"github.com/spf13/cobra"
)

type TodayCommand struct {
	log bool
}

func (cmd *TodayCommand) Cobra() *cobra.Command {
	cc := &cobra.Command{
		Use:   "today",
		Short: "what are the things you've done today?",
	}

	cc.Flags().BoolVarP(&cmd.log, "log", "l", false, "show time log, not summary")
	return cc
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

	if cmd.log {
		return cmd.printLog(entries)
	}

	toggl.PrintEntryList(entries)
	return nil
}

func (cmd *TodayCommand) printLog(entries []*toggl.Timer) error {
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
