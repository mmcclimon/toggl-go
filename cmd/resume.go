package cmd

import (
	"fmt"
	"time"

	t "github.com/mmmcclimon/toggl-go/internal/toggl"
	"github.com/spf13/cobra"
)

type ResumeCommand struct{}

func (cmd ResumeCommand) Cobra() *cobra.Command {
	return &cobra.Command{
		Use:   "resume",
		Short: "restart the last thing you were doing",
	}
}

func (cmd ResumeCommand) Run(toggl *t.Toggl, args []string) error {
	start := time.Now().Add(-6 * time.Hour)
	end := time.Now()

	entries, err := toggl.TimeEntries(start, end)
	if err != nil {
		return err
	}

	if len(entries) == 0 {
		fmt.Println("I dunno what you were last up to, sorry.")
		return nil
	}

	last := entries[0]

	timer, err := toggl.ResumeTimer(last)
	if err != nil {
		return err
	}

	fmt.Printf("started timer: %s\n", timer.OnelineDesc())
	return nil
}
