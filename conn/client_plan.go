package conn

import (
	"github.com/hashicorp/nomad/api"
)

// Plan is parsing the HCL code and registering the Job into Nomad
func (co *Client) Plan(code string, diff bool) (plan *api.JobPlanResponse, err error) {
	japi := co.c.Jobs()
	job, err := japi.ParseHCL(code, true)
	if err != nil {
		return
	}

	plan, _, err = japi.Plan(job, diff, nil)
	if err != nil {
		return nil, err
	}

	return plan, nil
}
