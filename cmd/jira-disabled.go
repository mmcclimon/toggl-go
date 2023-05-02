//go:build !jira

package cmd

import t "github.com/mmmcclimon/toggl-go/internal/toggl"

const JIRA_ENABLED = false

func startJiraTask(_ *t.Toggl, _ string) error {
	panic("ended in startJiraTask with no jira")
}
