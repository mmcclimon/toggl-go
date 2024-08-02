package toggl

import (
	"encoding/json"
	"io"
)

// We decode into a struct with all public members, then hide some of them publicly
type Project struct {
	Id     int
	Active bool
	Name   string
	Status string
}

func (c *Client) projectsFromResponseBody(body io.Reader) ([]Project, error) {
	var data []Project
	decoder := json.NewDecoder(body)

	if err := decoder.Decode(&data); err != nil {
		return nil, err
	}

	return data, nil
}
