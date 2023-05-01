package cmd

import (
	"fmt"
	"os"
	"regexp"
	"time"

	t "github.com/mmmcclimon/toggl-go/internal/toggl"
	"github.com/spf13/cobra"
)

var loggedLongHelp = "Search time entries for the last n days and find " +
	"descriptions matching DESC."

var loggedOpts struct {
	days int
}

var loggedCmd = &cobra.Command{
	Use:   "logged DESC",
	Short: "how much time did you spend on that thing?",
	Long:  loggedLongHelp,
	Run:   runLogged,
}

func init() {
	loggedCmd.Flags().IntVarP(&loggedOpts.days, "days", "d", 30, "how many days to look back")

	rootCmd.AddCommand(loggedCmd)
}

func runLogged(cmd *cobra.Command, args []string) {
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
	start := end.Add(-1 * time.Duration(loggedOpts.days) * 24 * time.Hour)

	entries, err := toggl.TimeEntries(start, end)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	matching := make([]*t.Timer, 0, len(entries))

	for _, entry := range entries {
		if re.MatchString(entry.OnelineDesc()) {
			matching = append(matching, entry)
		}
	}

	if len(matching) == 0 {
		fmt.Println("Nothing logged matching that description.")
		return
	}

	t.PrintEntryList(matching)
}
