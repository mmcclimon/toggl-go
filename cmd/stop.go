package cmd

import (
	"fmt"

	t "github.com/mmmcclimon/toggl-go/internal/toggl"
	"github.com/spf13/cobra"
)

type StopCommand struct{}

func (cmd StopCommand) Cobra() *cobra.Command {
	return &cobra.Command{
		Use:   "stop",
		Short: "stop doing the thing you're doing",
	}
}

func (cmd StopCommand) Run(toggl *t.Toggl, args []string) error {
	timer, err := toggl.StopCurrentTimer()

	if err != nil {
		switch err {
		case t.ErrNoTimer:
			fmt.Println("You don't have a running timer!")
			return nil
		default:
			return err
		}
	}

	fmt.Printf("spent %s: %s\n", timer.Duration(), timer.OnelineDesc())
	return nil
}
