package conn

import (
	"github.com/hashicorp/nomad/api"
)

type Client struct {
	c *api.Client
}

// NewClient returns a new client configured with the default values for Nomad
// See: https://godoc.org/github.com/hashicorp/nomad/api#DefaultConfig
func NewClient() (co *Client, err error) {
	co = &Client{}
	co.c, err = api.NewClient(api.DefaultConfig())
	if err != nil {
		return nil, err
	}

	return
}

// IsValid that is the question!
func (co *Client) IsValid(code string) bool {
	japi := co.c.Jobs()
	job, err := japi.ParseHCL(code, true)
	if err != nil {
		return false
	}

	res, _, err := japi.Validate(job, nil)
	if err != nil {
		return false
	}

	return len(res.Error) == 0 && len(res.Warnings) == 0 && len(res.ValidationErrors) == 0
}
