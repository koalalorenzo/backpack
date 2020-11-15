package conn

import (
	"github.com/hashicorp/nomad/api"
)

type Client struct {
	c     *api.Client
	jobs  *api.Jobs
	alloc *api.Allocations
}

// NewClient returns a new client configured with the default values for Nomad
// See: https://godoc.org/github.com/hashicorp/nomad/api#DefaultConfig
func NewClient() (co *Client, err error) {
	co = &Client{}
	co.c, err = api.NewClient(api.DefaultConfig())
	if err != nil {
		return nil, err
	}
	co.jobs = co.c.Jobs()
	co.alloc = co.c.Allocations()
	return
}

// IsValid that is the question!
func (co *Client) IsValid(code string) bool {
	job, err := co.GetJobFromCode(code)
	if err != nil {
		return false
	}

	res, _, err := co.jobs.Validate(job, nil)
	if err != nil {
		return false
	}

	return len(res.Error) == 0 && len(res.Warnings) == 0 && len(res.ValidationErrors) == 0
}
