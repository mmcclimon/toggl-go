package cmd

import (
	"fmt"

	t "github.com/mmmcclimon/toggl-go/internal/toggl"
	"github.com/spf13/cobra"
)

type TimerCommand struct{}

func (cmd TimerCommand) Cobra() *cobra.Command {
	return &cobra.Command{
		Use:   "timer",
		Short: "what are you doing right now?",
	}
}

func (cmd TimerCommand) Run(toggl *t.Toggl, args []string) error {
	timer, err := toggl.CurrentTimer()

	if err != nil {
		switch err {
		case t.ErrNoTimer:
			fmt.Println("You don't have a running timer!")
			return nil
		default:
			return err
		}
	}

	fmt.Printf("%s so far: %s\n", timer.Duration(), timer.OnelineDesc())
	return nil
}
