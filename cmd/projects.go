package cmd

import (
	"fmt"
	"sort"

	"github.com/mmmcclimon/toggl-go/internal/toggl"
	"github.com/spf13/cobra"
	"golang.org/x/exp/maps"
)

type ProjectsCommand struct{}

func (cmd ProjectsCommand) Cobra() *cobra.Command {
	return &cobra.Command{
		Use:   "projects",
		Short: "list the buckets things can go in",
	}
}

func (cmd ProjectsCommand) Run(tc *toggl.Client, args []string) error {
	projects := tc.Config.ProjectShortcuts

	shortcuts := maps.Keys(projects)
	sort.Strings(shortcuts)

	for _, sc := range shortcuts {
		fmt.Printf("- %s (%d)\n", sc, projects[sc])
	}

	return nil
}
