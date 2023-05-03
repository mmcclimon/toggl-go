package toggl

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/mmmcclimon/toggl-go/internal/jira"
)

type Config struct {
	ApiToken         string                       `toml:"api_token"`
	WorkspaceId      int                          `toml:"workspace_id"`
	ProjectShortcuts map[string]int               `toml:"project_shortcuts"`
	TaskShortcuts    map[string]map[string]string `toml:"task_shortcuts"`
	JiraConfig       *jira.Config                 `toml:"jira"`
	projectsById     map[int]string
}

// ReadConfig reads the toggl config file, and returns an error if it can't
// figure out what to read, or if it's not toml
func (c *Client) ReadConfig() error {
	filename := os.Getenv("TOGGL_CONFIG_FILE")

	if filename == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("could not determine homedir: %w", err)
		}

		filename = filepath.Join(home, ".togglrc")
	}

	tomlData, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("could not read config file: %w", err)
	}

	_, err = toml.Decode(string(tomlData), &c.Config)
	if err != nil {
		return fmt.Errorf("bad config file: %w", err)
	}

	byId := make(map[int]string)
	for name, id := range c.Config.ProjectShortcuts {
		byId[id] = name
	}

	c.Config.projectsById = byId

	return nil
}
