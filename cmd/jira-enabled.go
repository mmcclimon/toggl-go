//go:build jira

package cmd

import "github.com/mmmcclimon/toggl-go/internal/toggl"

const JIRA_ENABLED = true

// used by cmd/start
func startJiraTask(tc *toggl.Client, taskId string, projectId int) error {
	c := tc.Config.NewJiraClient()
	issue := c.GetIssue(taskId)

	if projectId == 0 {
		projectId = issue.TogglProjectId
	}

	return startTask(tc, issue.PrettyDescription(), projectId, "")
}
