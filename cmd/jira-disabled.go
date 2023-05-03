//go:build !jira

package cmd

import "github.com/mmmcclimon/toggl-go/internal/toggl"

const JIRA_ENABLED = false

func startJiraTask(_ *toggl.Client, _ string) error {
	panic("ended in startJiraTask with no jira")
}
