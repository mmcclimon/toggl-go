package cmd

import (
	"fmt"
	"os"
	"regexp"
	"time"

	"github.com/mmmcclimon/toggl-go/internal/toggl"
	"github.com/spf13/cobra"
)

type LoggedCommand struct {
	days int
}

func (cmd LoggedCommand) Cobra() *cobra.Command {
	long := "Search time entries for the last n days and find descriptions matching DESC."

	cc := &cobra.Command{
		Use:   "logged DESC",
		Short: "how much time did you spend on that thing?",
		Long:  long,
	}

	cc.Flags().IntVarP(&cmd.days, "days", "d", 30, "how many days to look back")
	return cc
}

func (cmd LoggedCommand) Run(tc *toggl.Client, args []string) error {
	if len(args) != 1 {
		fmt.Println("need exactly one description to search for")
		os.Exit(1)
	}

	re, err := regexp.Compile("(?i)" + args[0])
	if err != nil {
		fmt.Println("could not compile description regex:", err)
		os.Exit(1)
	}

	end := time.Now()
	start := end.Add(-1 * time.Duration(cmd.days) * 24 * time.Hour)

	entries, err := tc.TimeEntries(start, end)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	matching := make([]*toggl.Timer, 0, len(entries))

	for _, entry := range entries {
		if re.MatchString(entry.OnelineDesc()) {
			matching = append(matching, entry)
		}
	}

	if len(matching) == 0 {
		fmt.Println("Nothing logged matching that description.")
		return nil
	}

	toggl.PrintEntryList(matching)
	return nil
}
