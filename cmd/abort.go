package cmd

import (
	"fmt"

	"github.com/mmmcclimon/toggl-go/internal/toggl"
	"github.com/spf13/cobra"
)

type AbortCommand struct{}

func (cmd *AbortCommand) Cobra() *cobra.Command {
	return &cobra.Command{
		Use:   "abort",
		Short: "actually, you weren't doing that thing after all",
	}
}

func (cmd *AbortCommand) Run(tc *toggl.Client, args []string) error {
	timer, err := tc.AbortCurrentTimer()

	if err != nil {
		switch err {
		case toggl.ErrNoTimer:
			fmt.Println("You don't have a running timer!")
			return nil
		default:
			return err
		}
	}

	fmt.Printf("aborted timer: %s\n", timer.OnelineDesc())
	return nil
}
