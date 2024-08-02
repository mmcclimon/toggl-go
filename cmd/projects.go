package cmd

import (
	"fmt"
	"sort"
	"strings"

	"github.com/mmmcclimon/toggl-go/internal/toggl"
	"github.com/spf13/cobra"
	"golang.org/x/exp/maps"
)

type ProjectsCommand struct {
	remote bool
}

func (cmd *ProjectsCommand) Cobra() *cobra.Command {
	cc := &cobra.Command{
		Use:   "projects",
		Short: "list the buckets things can go in",
	}

	cc.Flags().BoolVarP(&cmd.remote, "remote", "r", false, "fetch remote projects")
	return cc
}

func (cmd *ProjectsCommand) Run(tc *toggl.Client, args []string) error {
	if cmd.remote {
		return cmd.listRemoteProjects(tc)
	}

	projects := tc.Config.ProjectShortcuts

	shortcuts := maps.Keys(projects)
	sort.Strings(shortcuts)

	for _, sc := range shortcuts {
		fmt.Printf("- %s (%d)\n", sc, projects[sc])
	}

	return nil
}

func (cmd *ProjectsCommand) listRemoteProjects(tc *toggl.Client) error {
	projects, err := tc.Projects()
	if err != nil {
		return err
	}

	wrap := func(s string) string {
		if strings.Contains(s, " ") {
			return fmt.Sprintf(`"%s"`, s)
		}

		return s
	}

	for _, p := range projects {
		fmt.Printf("%s = %d\n", wrap(p.Name), p.Id)
	}

	return nil
}
