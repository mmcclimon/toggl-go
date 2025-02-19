package cmd

import (
	"fmt"
	"maps"
	"slices"

	"github.com/mmmcclimon/toggl-go/internal/toggl"
	"github.com/spf13/cobra"
)

type ShortcutsCommand struct{}

func (cmd *ShortcutsCommand) Cobra() *cobra.Command {
	return &cobra.Command{
		Use:   "shortcuts",
		Short: "list the things you can start easily",
	}
}

func (cmd *ShortcutsCommand) Run(tc *toggl.Client, args []string) error {
	shortcuts := tc.Config.TaskShortcuts

	titles := slices.Sorted(maps.Keys(shortcuts))

	length := 0
	for _, title := range titles {
		if len(title) > length {
			length = len(title)
		}
	}

	for _, title := range titles {
		shortcut := shortcuts[title]
		desc := shortcut["desc"]
		project := shortcut["project"]

		// descriptionless task, just use the project as the description
		if desc == "" {
			desc = project
			project = "*taskless*"
		}

		fmt.Printf("@%-*s %s (%s)\n", length+2, title, desc, project)
	}

	return nil
}
