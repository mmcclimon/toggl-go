package cmd

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	t "github.com/mmmcclimon/toggl-go/internal/toggl"
	"github.com/spf13/cobra"
)

type StartCommand struct {
	project string
	id      string
}

func (cmd StartCommand) Cobra() *cobra.Command {
	cc := &cobra.Command{
		Use:   "start description",
		Short: "start doing a new thing",
	}

	cc.Flags().StringVarP(&cmd.project, "project", "p", "", "project shortcut for this task")

	if JIRA_ENABLED {
		cc.Flags().StringVarP(&cmd.id, "id", "i", "", "jira id for this task")
	}

	return cc
}

func (cmd StartCommand) Run(toggl *t.Toggl, args []string) error {
	desc := strings.Join(args, " ")
	if len(desc) == 0 {
		return errors.New("need a description")
	}

	likelyId := regexp.MustCompile(`(?i)^[a-z]{3,}-[0-9]+$`)

	if JIRA_ENABLED && (cmd.id != "" || likelyId.MatchString(desc)) {
		id := cmd.id
		if id == "" {
			id = desc
		}

		return startJiraTask(toggl, id)
	}

	projectId := 0
	if len(cmd.project) > 0 {
		projectId = toggl.Config.ProjectShortcuts[cmd.project]
	}

	// is this a shortcut
	if strings.HasPrefix(desc, "@") {
		fields := strings.Fields(desc)
		sc := fields[0]

		shortcut, ok := toggl.Config.TaskShortcuts[strings.TrimPrefix(sc, "@")]
		if !ok {
			return fmt.Errorf("could not resolve shortcut %s", sc)
		}

		// no error handling here, just don't mess up your config file, ok
		desc = shortcut["desc"]
		if proj, ok := shortcut["project"]; ok {
			projectId = toggl.Config.ProjectShortcuts[proj]
		}

		// add back on the tag, if there is one
		if len(fields) > 1 {
			desc += " " + strings.Join(fields[1:], " ")
		}
	}

	// tags
	tag := ""
	words := strings.Split(desc, " ")
	last := words[len(words)-1]

	if strings.HasPrefix(last, "#") {
		tag = strings.TrimPrefix(last, "#")
		desc = strings.Join(words[0:len(words)-1], " ")
	}

	return startTask(toggl, desc, projectId, tag)
}

func startTask(toggl *t.Toggl, desc string, projectId int, tag string) error {
	timer, err := toggl.StartTimer(desc, projectId, tag)
	if err != nil {
		return err
	}

	fmt.Printf("started timer: %s\n", timer.OnelineDesc())
	return nil
}
