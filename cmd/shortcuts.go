package cmd

import (
	"fmt"
	"sort"

	t "github.com/mmmcclimon/toggl-go/internal/toggl"
	"github.com/spf13/cobra"
	"golang.org/x/exp/maps"
)

type ShortcutsCommand struct{}

func (cmd ShortcutsCommand) Cobra() *cobra.Command {
	return &cobra.Command{
		Use:   "shortcuts",
		Short: "list the things you can start easily",
	}
}

func (cmd ShortcutsCommand) Run(toggl *t.Toggl, args []string) error {
	shortcuts := toggl.Config.TaskShortcuts

	titles := maps.Keys(shortcuts)
	sort.Strings(titles)

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
