//go:build jira

package cmd

import t "github.com/mmmcclimon/toggl-go/internal/toggl"

const JIRA_ENABLED = true

// used by cmd/start
func startJiraTask(toggl *t.Toggl, taskId string) error {
	c := toggl.Config.NewJiraClient()
	issue := c.GetIssue(taskId)
	return startTask(toggl, issue.PrettyDescription(), issue.TogglProjectId, "")
}
