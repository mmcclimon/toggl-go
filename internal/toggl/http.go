package toggl

import (
	"net/http"
	"net/url"
	"os"
)

func (c *Client) doRequest(req *http.Request) (*http.Response, error) {
	// default headers
	req.Header.Add("Accept", "application/json")
	req.SetBasicAuth(c.Config.ApiToken, "api_token")

	if req.Body != nil {
		req.Header.Add("Content-Type", "application/json")
	}

	res, err := c.client.Do(req)

	if err != nil {
		return nil, err
	} else if res.StatusCode >= 400 {
		// just bail, I'm never going to do anything useful with this.
		dumpResponseAndExit(res)
	}

	return res, nil
}

func (c *Client) get(loc *url.URL) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, loc.String(), nil)
	if err != nil {
		panic(err) // should not happen
	}

	return c.doRequest(req)
}

func dumpResponseAndExit(res *http.Response) {
	err := res.Write(os.Stderr)
	if err != nil {
		panic("could not write out response to stderr")
	}

	os.Exit(1)
}
