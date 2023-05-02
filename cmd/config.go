package cmd

import (
	"github.com/davecgh/go-spew/spew"
	t "github.com/mmmcclimon/toggl-go/internal/toggl"
	"github.com/spf13/cobra"
)

type ConfigCommand struct{}

func (cmd ConfigCommand) Cobra() *cobra.Command {
	return &cobra.Command{
		Use:    "config",
		Short:  "dump config and exit (for debugging)",
		Hidden: true,
	}
}

func (cmd ConfigCommand) Run(toggl *t.Toggl, args []string) error {
	cfg := toggl.Config
	spew.Dump(cfg)
	return nil
}
